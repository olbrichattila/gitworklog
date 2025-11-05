package dto

import (
	"time"
)

type Config struct {
	Username     string `yaml:"username"`
	Repositories []struct {
		Path string `yaml:"path"`
	} `yaml:"repositories"`
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
