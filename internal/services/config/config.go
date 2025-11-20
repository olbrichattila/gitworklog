package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/dto"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"

	yamlReader "gopkg.in/yaml.v3"
)

const (
	configDirectory = ".gitworklog"
	configFileName  = "config.yaml"
)

func New() (contracts.ConfigProvider, error) {
	var err error
	s := &service{}
	s.fileName, err = s.configFileName()
	if err != nil {
		return nil, err
	}

	return s, err
}

type service struct {
	fileName string
}

func (s *service) AddRepository(fullPath string) error {
	config, err := s.Get()
	if err != nil {
		return err
	}

	config.Repositories = append(config.Repositories, dto.Repository{Path: fullPath})

	return s.save(s.fileName, config)
}

func (s *service) GetRepositories() ([]string, error) {
	config, err := s.Get()
	if err != nil {
		return nil, err
	}

	result := make([]string, len(config.Repositories))
	for i, repo := range config.Repositories {
		result[i] = repo.Path
	}

	return result, nil
}

func (s *service) GetUserName() (string, error) {
	config, err := s.Get()
	if err != nil {
		return "", err
	}

	return config.Username, nil
}

func (s *service) RemoveRepository(fullPath string) error {
	config, err := s.Get()
	if err != nil {
		return err
	}

	i, ok := s.getElementByIndex(config.Repositories, fullPath)
	if !ok {
		return worklogerrors.Wrap(worklogerrors.ErrGitRepositoryPathNotFound, err, fullPath)
	}

	config.Repositories = s.removeByIndex(config.Repositories, i)

	return s.save(s.fileName, config)
}

func (s *service) SetUserName(name string) error {
	config, err := s.Get()
	if err != nil {
		return err
	}

	config.Username = name

	return s.save(s.fileName, config)
}

func (s *service) Get() (dto.Config, error) {
	data, err := os.ReadFile(s.fileName)
	if err != nil {
		return dto.Config{}, worklogerrors.Wrap(worklogerrors.ErrConfigFileRead, err, s.fileName)
	}

	var params dto.Config
	if err := yamlReader.Unmarshal(data, &params); err != nil {
		return dto.Config{}, fmt.Errorf("error parsing YAML: %w", err)
	}

	return params, nil
}

func (s *service) save(fileName string, config dto.Config) error {
	out, err := yamlReader.Marshal(config)
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrParseConfig, err, s.fileName)
	}

	err = os.WriteFile(fileName, out, 0755)
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrConfigFileRead, err, s.fileName)
	}

	return nil
}

func (s *service) configFileName() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", worklogerrors.Wrap(worklogerrors.ErrGetHomeDirectory, err, "")
	}

	dir := filepath.Join(home, configDirectory)
	config := filepath.Join(dir, configFileName)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", worklogerrors.Wrap(worklogerrors.ErrCreatingConfigDirectory, err, dir)
	}

	if _, err := os.Stat(config); os.IsNotExist(err) {
		err := s.save(config, dto.Config{})
		if err != nil {
			return "", err
		}
	}

	return config, nil
}

func (s *service) removeByIndex(gitRepoNameList []dto.Repository, i int) []dto.Repository {
	if i < 0 || i >= len(gitRepoNameList) {
		return gitRepoNameList
	}
	return append(gitRepoNameList[:i], gitRepoNameList[i+1:]...)
}

func (s *service) getElementByIndex(gitRepoNameList []dto.Repository, path string) (int, bool) {
	for i, repo := range gitRepoNameList {
		if repo.Path == path {
			return i, true
		}
	}

	return -1, false
}
