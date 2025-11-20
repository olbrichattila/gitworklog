package worklogerrors

import (
	"errors"
	"fmt"
)

var (
	ErrNilValue                              = errors.New("nil value")
	ErrApplication                           = errors.New("application error")
	ErrIncorrectNumberOfParameters           = errors.New("incorrect number of parameters")
	ErrIncorrectDateFormat                   = errors.New("incorrect date format")
	ErrConfigFileMissing                     = errors.New("missing config file")
	ErrConfigFileRead                        = errors.New("cannot read config file")
	ErrConfigFileWrite                       = errors.New("cannot write config file")
	ErrCannotOpenGitRepo                     = errors.New("cannot open git repository")
	ErrNoRepositories                        = errors.New("no repository paths are configured")
	ErrMissingUserName                       = errors.New("user name not configured")
	ErrFailedToGetGitBranches                = errors.New("failed to get git branches")
	ErrIteratingGitBranches                  = errors.New("cannot iterate git branches")
	ErrAggregation                           = errors.New("cannot aggregate data")
	ErrParseConfig                           = errors.New("error parsing config")
	ErrGetHomeDirectory                      = errors.New("error getting home directory")
	ErrCreatingConfigDirectory               = errors.New("error creating config directory")
	ErrNotEnoughCommandLineParameter         = errors.New("not enough command line parameter")
	ErrIncorrectCommandLineParameter         = errors.New("incorrect command line parameter")
	ErrIncorrectNumberOfCommandLineParameter = errors.New("incorrect number of command line parameter")
	ErrCmdConfigProvider                     = errors.New("command line config provider error")
	ErrGitRepositoryPathNotFound             = errors.New("git repository path not found")
)

// Wrap adds context to a parent error while preserving error wrapping.
func Wrap(base, parent error, context string) error {
	if parent != nil {
		return fmt.Errorf("%s: %v: %w", context, base, parent)
	}

	return fmt.Errorf("%s: %w", context, base)
}
