package worklogerrors

import (
	"errors"
	"fmt"
)

var (
	ErrNilValue                    = errors.New("nil value")
	ErrApplication                 = errors.New("application error")
	ErrIncorrectNumberOfParameters = errors.New("incorrect number of parameters")
	ErrIncorrectDateFormat         = errors.New("incorrect date format")
	ErrConfigFileMissing           = errors.New("missing config file")
	ErrConfigFileRead              = errors.New("cannot read config file")
	ErrCannotOpenGitRepo           = errors.New("cannot open git repository")
	ErrFailedToGetGitBranches      = errors.New("failed to get git branches")
	ErrIteratingGitBranches        = errors.New("cannot iterate git branches")
	ErrAggregation                 = errors.New("cannot aggregate data")
)

// Wrap adds context to a parent error while preserving error wrapping.
func Wrap(base, parent error, context string) error {
	if parent != nil {
		return fmt.Errorf("%s: %v: %w", context, base, parent)
	}

	return fmt.Errorf("%s: %w", context, base)
}
