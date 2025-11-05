package reportdisplay

import (
	"fmt"
	"sort"
	"time"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/dto"
)

func New() contracts.ReportDisplay {
	return &service{}
}

type service struct {
}

func (s *service) Display(commitDateString string, aggregate map[string][]dto.AggregateItem) {
	fmt.Printf("====%s====\n", commitDateString)

	aggregatedItems := aggregate[commitDateString]
	sort.Slice(aggregatedItems, func(i, j int) bool {
		return aggregatedItems[i].GitCommit.DateTime.Before(aggregatedItems[j].GitCommit.DateTime)
	})

	for _, commit := range aggregatedItems {
		fmt.Printf(
			"------------------\n - Repo: %s\n - Branch: %s\n - Message: %s\n - Time: %s\n",
			commit.RepoName,
			commit.GitCommit.BranchName,
			commit.GitCommit.CommitMessage,
			commit.GitCommit.DateTime.Format(time.TimeOnly),
		)
	}
	fmt.Println()
}
