package app

import (
	"sort"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/dto"
	"github.com/olbrichattila/gitworklog/internal/services/cmdparams"
	"github.com/olbrichattila/gitworklog/internal/services/config"
	"github.com/olbrichattila/gitworklog/internal/services/gitmanager"
	"github.com/olbrichattila/gitworklog/internal/services/reportaggregator"
	"github.com/olbrichattila/gitworklog/internal/services/reportdisplay"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"
)

const (
	configFileName = "./config.yaml"
)

type application struct {
	GitManager       contracts.GitManager
	ConfigProvider   contracts.ConfigProvider
	CmdParamProvider contracts.CmdParamProvider
	ReportAggregator contracts.ReportAggregator
	ReportDisplay    contracts.ReportDisplay
}

func New() (contracts.AppProvider, error) {
	cmdPars := cmdparams.New()
	gitMgr := gitmanager.New()
	reportDisplay := reportdisplay.New()

	cnf, err := config.New(configFileName)
	if err != nil {
		return nil, worklogerrors.Wrap(worklogerrors.ErrApplication, err, "config")
	}

	reportAgg, err := reportaggregator.New(
		gitMgr,
	)
	if err != nil {
		return nil, worklogerrors.Wrap(worklogerrors.ErrApplication, err, "config")
	}

	app := &application{
		GitManager:       gitMgr,
		ConfigProvider:   cnf,
		CmdParamProvider: cmdPars,
		ReportAggregator: reportAgg,
		ReportDisplay:    reportDisplay,
	}

	return app, nil
}

func (a *application) Run() error {
	cmdParams, err := a.CmdParamProvider.Get()
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrApplication, err, "cmd params")
	}

	configValues, err := a.ConfigProvider.Get()
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrApplication, err, "config")
	}

	aggregate, err := a.ReportAggregator.Aggregate(configValues, cmdParams)
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrApplication, err, "aggregation")
	}

	keys := a.sortedKeys(aggregate)
	for _, commitDateString := range keys {
		a.ReportDisplay.Display(commitDateString, aggregate)
	}

	return nil
}

func (a *application) sortedKeys(aggregate map[string][]dto.AggregateItem) []string {
	keys := make([]string, 0, len(aggregate))
	for k := range aggregate {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
