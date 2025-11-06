package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/olbrichattila/gitworklog/internal/app"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "help" {
		displayUsage()
		return
	}

	app, err := app.New()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = app.Run(); err != nil {
		if errors.Is(err, worklogerrors.ErrIncorrectNumberOfParameters) {
			displayUsage()
			return
		}
		fmt.Println(err.Error())
	}
}

func displayUsage() {
	fmt.Println(
		`Usage:
gitworklog <?fromDate> <?toDate>

Example:
gitworklog 2025-06-05 2025-06-10

For single date:
gitworklog 2025-06-05

For today:
gitworklog 2025-06-05

This help:
gitworklog help`,
	)
}
