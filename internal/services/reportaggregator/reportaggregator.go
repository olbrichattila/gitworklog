package reportaggregator

import (
	"time"

	"path/filepath"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/dto"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"
)

func New(gitManagerService contracts.GitManager) (contracts.ReportAggregator, error) {
	if gitManagerService == nil {
		return nil, worklogerrors.Wrap(worklogerrors.ErrNilValue, nil, "reportAggregator, gitManager")
	}

	return &service{
		gitManager: gitManagerService,
	}, nil
}

type service struct {
	gitManager contracts.GitManager
}

func (s *service) Aggregate(configValues dto.Config, cmdParams dto.CmdParams) (map[string][]dto.AggregateItem, error) {
	aggregate := map[string][]dto.AggregateItem{}
	if configValues.Username == "" {
		return aggregate, worklogerrors.Wrap(worklogerrors.ErrMissingUserName, nil, "config")
	}

	if len(configValues.Repositories) == 0 {
		return aggregate, worklogerrors.Wrap(worklogerrors.ErrNoRepositories, nil, "config")
	}
	for _, gitRepos := range configValues.Repositories {
		res, err := s.gitManager.Log(gitRepos.Path, configValues.Username, cmdParams.From, cmdParams.To)
		if err != nil {
			return nil, worklogerrors.Wrap(worklogerrors.ErrAggregation, err, gitRepos.Path)
		}

		for _, commit := range res {
			key := commit.DateTime.Format(time.DateOnly)
			aggregate[key] = append(
				aggregate[key],
				dto.AggregateItem{
					RepoName:  filepath.Base(gitRepos.Path),
					GitCommit: commit,
				},
			)
		}
	}
	return aggregate, nil
}
