package contracts

import (
	"time"

	"github.com/olbrichattila/gitworklog/internal/dto"
)

type GitManager interface {
	Log(repoPath, authorEmail string, from time.Time, to time.Time) ([]dto.GitCommit, error)
}

type ConfigProvider interface {
	Get() (dto.Config, error)
}

type CmdParamProvider interface {
	Get() (dto.CmdParams, error)
}

type ReportAggregator interface {
	Aggregate(configValues dto.Config, cmdParams dto.CmdParams) (map[string][]dto.AggregateItem, error)
}

type ReportDisplay interface {
	Display(commitDateString string, aggregate map[string][]dto.AggregateItem)
}
