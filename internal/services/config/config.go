package config

import (
	"fmt"
	"os"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/dto"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"

	yamlReader "gopkg.in/yaml.v3"
)

func New(fileName string) (contracts.ConfigProvider, error) {
	if fileName == "" {
		return nil, worklogerrors.Wrap(worklogerrors.ErrConfigFileMissing, nil, "")
	}

	return &service{
		fileName: fileName,
	}, nil
}

type service struct {
	fileName string
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
