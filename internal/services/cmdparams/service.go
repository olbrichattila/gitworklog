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
	parCnt := len(os.Args)
	switch parCnt {
	case 1:
		return s.getToday()
	case 2:
		return s.getSingleDate()
	case 3:
		return s.getDateRange()
	default:
		return dto.CmdParams{}, worklogerrors.Wrap(worklogerrors.ErrIncorrectNumberOfParameters, nil, strconv.Itoa(parCnt))
	}
}

func (s *service) getToday() (dto.CmdParams, error) {
	t := time.Now()
	return dto.CmdParams{
		From: time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()),
		To:   time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location()),
	}, nil
}

func (s *service) getSingleDate() (dto.CmdParams, error) {
	t, err := time.Parse(time.DateOnly, os.Args[1])
	if err != nil {
		return dto.CmdParams{}, worklogerrors.Wrap(worklogerrors.ErrIncorrectDateFormat, nil, os.Args[1])
	}
	return dto.CmdParams{
		From: t,
		To:   time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location()),
	}, nil
}

func (s *service) getDateRange() (dto.CmdParams, error) {
	from, err := time.Parse(time.DateOnly, os.Args[1])
	if err != nil {
		return dto.CmdParams{}, worklogerrors.Wrap(worklogerrors.ErrIncorrectDateFormat, nil, os.Args[1])
	}

	to, err := time.Parse(time.DateOnly, os.Args[2])
	if err != nil {
		return dto.CmdParams{}, worklogerrors.Wrap(worklogerrors.ErrIncorrectDateFormat, nil, os.Args[2])
	}

	return dto.CmdParams{
		From: from,
		To:   time.Date(to.Year(), to.Month(), to.Day(), 23, 59, 59, 0, to.Location()),
	}, nil
}
