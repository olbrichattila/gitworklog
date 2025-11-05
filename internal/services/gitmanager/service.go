package gitmanager

import (
	"strings"
	"time"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/dto"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func New() contracts.GitManager {
	return &service{}
}

type service struct {
}

func (s *service) Log(repoPath, authorEmail string, from time.Time, to time.Time) ([]dto.GitCommit, error) {
	result := []dto.GitCommit{}

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, worklogerrors.Wrap(worklogerrors.ErrCannotOpenGitRepo, err, repoPath)
	}

	branches, err := repo.Branches()
	if err != nil {
		return nil, worklogerrors.Wrap(worklogerrors.ErrFailedToGetGitBranches, err, repoPath)
	}

	err = branches.ForEach(func(ref *plumbing.Reference) error {
		branchName := ref.Name().Short()
		commitIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
		if err != nil {
			return err
		}

		return commitIter.ForEach(func(c *object.Commit) error {
			if c.Author.When.After(from) && c.Author.When.Before(to) && c.Author.Email == authorEmail {
				result = append(
					result,
					dto.GitCommit{
						DateTime:      c.Author.When,
						BranchName:    branchName,
						CommitMessage: strings.TrimSpace(c.Message),
					},
				)
			}
			return nil
		})
	})
	if err != nil {
		return nil, worklogerrors.Wrap(worklogerrors.ErrIteratingGitBranches, err, repoPath)
	}

	return result, nil
}
