package configcmd

import (
	"fmt"
	"os"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"
)

const (
	argUserName = "set-name"
	argListName = "get-name"
	argAddRepo  = "add-repository"
	argDelRepo  = "delete-repository"
	argListRepo = "list-repositories"
)

func New(configProvider contracts.ConfigProvider) (contracts.ConfigCmdProvider, error) {
	if configProvider == nil {
		return nil, worklogerrors.Wrap(worklogerrors.ErrNilValue, nil, "configCmdProvider")
	}

	return &service{
		configProvider: configProvider,
	}, nil
}

type service struct {
	configProvider contracts.ConfigProvider
}

func (s *service) Run() (bool, error) {
	if len(os.Args) == 1 {
		return false, nil
	}

	firstArg := os.Args[1]
	if firstArg == "config" {
		return true, s.callSelectedConfig()
	}

	return false, nil
}

func (s *service) callSelectedConfig() error {
	if len(os.Args) < 3 {
		return worklogerrors.Wrap(worklogerrors.ErrNotEnoughCommandLineParameter, nil, "configCmdProvider")
	}

	switch os.Args[2] {
	case argUserName:
		return s.editUserName()
	case argAddRepo:
		return s.addRepo()
	case argDelRepo:
		return s.deleteRepo()
	case argListRepo:
		return s.displayRepoList()
	case argListName:
		return s.displayName()
	}

	return worklogerrors.Wrap(worklogerrors.ErrIncorrectCommandLineParameter, nil, os.Args[2])
}

func (s *service) editUserName() error {
	if len(os.Args) != 4 {
		return worklogerrors.Wrap(worklogerrors.ErrIncorrectNumberOfCommandLineParameter, nil, argUserName)
	}

	return s.configProvider.SetUserName(os.Args[3])
}

func (s *service) addRepo() error {
	if len(os.Args) != 4 {
		return worklogerrors.Wrap(worklogerrors.ErrIncorrectNumberOfCommandLineParameter, nil, argUserName)
	}

	return s.configProvider.AddRepository(os.Args[3])
}

func (s *service) deleteRepo() error {
	if len(os.Args) != 4 {
		return worklogerrors.Wrap(worklogerrors.ErrIncorrectNumberOfCommandLineParameter, nil, argUserName)
	}

	return s.configProvider.RemoveRepository(os.Args[3])
}

func (s *service) displayRepoList() error {
	repoList, err := s.configProvider.GetRepositories()
	if err != nil {
		return err
	}

	for _, repoName := range repoList {
		fmt.Println(repoName)
	}

	return nil
}

func (s *service) displayName() error {
	userName, err := s.configProvider.GetUserName()
	if err != nil {
		return err
	}

	fmt.Println(userName)

	return nil
}
