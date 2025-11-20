package app

import (
	"sort"

	"github.com/olbrichattila/gitworklog/internal/contracts"
	"github.com/olbrichattila/gitworklog/internal/dto"
	"github.com/olbrichattila/gitworklog/internal/services/cmdparams"
	"github.com/olbrichattila/gitworklog/internal/services/config"
	"github.com/olbrichattila/gitworklog/internal/services/configcmd"
	"github.com/olbrichattila/gitworklog/internal/services/gitmanager"
	"github.com/olbrichattila/gitworklog/internal/services/reportaggregator"
	"github.com/olbrichattila/gitworklog/internal/services/reportdisplay"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"
)

type application struct {
	gitManager        contracts.GitManager
	configProvider    contracts.ConfigProvider
	configCmdProvider contracts.ConfigCmdProvider
	cmdParamProvider  contracts.CmdParamProvider
	reportAggregator  contracts.ReportAggregator
	reportDisplay     contracts.ReportDisplay
}

func New() (contracts.AppProvider, error) {
	cmdPars := cmdparams.New()

	gitMgr := gitmanager.New()
	reportDisplay := reportdisplay.New()

	cnf, err := config.New()
	if err != nil {
		return nil, worklogerrors.Wrap(worklogerrors.ErrApplication, err, "config")
	}

	configCmdPars, err := configcmd.New(cnf)
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
		gitManager:        gitMgr,
		configProvider:    cnf,
		cmdParamProvider:  cmdPars,
		reportAggregator:  reportAgg,
		reportDisplay:     reportDisplay,
		configCmdProvider: configCmdPars,
	}

	return app, nil
}

func (a *application) Run() error {
	executed, err := a.configCmdProvider.Run()
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrApplication, err, "cmd config parameters")
	}

	// If it was only configuration, exit
	if executed {
		return nil
	}

	cmdParams, err := a.cmdParamProvider.Get()
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrApplication, err, "cmd params")
	}

	configValues, err := a.configProvider.Get()
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrApplication, err, "config")
	}

	aggregate, err := a.reportAggregator.Aggregate(configValues, cmdParams)
	if err != nil {
		return worklogerrors.Wrap(worklogerrors.ErrApplication, err, "aggregation")
	}

	keys := a.sortedKeys(aggregate)
	for _, commitDateString := range keys {
		a.reportDisplay.Display(commitDateString, aggregate)
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
