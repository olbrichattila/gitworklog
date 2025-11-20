package dto

import (
	"time"
)

type Repository struct {
	Path string `yaml:"path"`
}

type Config struct {
	Username     string       `yaml:"username"`
	Repositories []Repository `yaml:"repositories"`
}

type GitCommit struct {
	DateTime      time.Time
	BranchName    string
	CommitMessage string
}

type CmdParams struct {
	From time.Time
	To   time.Time
}

type AggregateItem struct {
	RepoName  string
	GitCommit GitCommit
}
