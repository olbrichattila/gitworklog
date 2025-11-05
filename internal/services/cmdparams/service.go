package cmdparams

import (
	"os"
	"strconv"
	"time"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/dto"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"
)

func New() contracts.CmdParamProvider {
	return &service{}
}

type service struct {
}

func (s *service) Get() (dto.CmdParams, error) {
	var err error
	result := dto.CmdParams{}
	if len(os.Args) < 2 {
		return result, worklogerrors.Wrap(worklogerrors.ErrIncorrectNumberOfParameters, nil, strconv.Itoa(len(os.Args)-1))
	}

	result.From, err = time.Parse(time.DateOnly, os.Args[1])
	if err != nil {
		return result, worklogerrors.Wrap(worklogerrors.ErrIncorrectDateFormat, nil, os.Args[1])
	}

	if len(os.Args) == 2 {
		result.To = time.Date(result.From.Year(), result.From.Month(), result.From.Day(), 23, 59, 59, 0, result.From.Location())
		return result, nil
	}

	toDate, err := time.Parse(time.DateOnly, os.Args[2])
	if err != nil {
		return result, worklogerrors.Wrap(worklogerrors.ErrIncorrectDateFormat, nil, os.Args[2])
	}

	result.To = time.Date(toDate.Year(), toDate.Month(), toDate.Day(), 23, 59, 59, 0, toDate.Location())

	return result, nil
}
