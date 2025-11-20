package main

import (
	"errors"
	"fmt"

	"github.com/olbrichattila/gitworklog/internal/app"
	"github.com/olbrichattila/gitworklog/internal/worklogerrors"
)

func main() {
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
gitworklog <fromDate> [toDate]

Example:
gitworklog 2025-06-05 2025-06-10

For single date:
gitworklog 2025-06-05

For today:
gitworklog 2025-06-05

Configuration:
You can configure your command with the following commands:

Add a new git user email:
	gitworklog config set-name <user email address>
Add a new local repository path:	
	gitworklog config add-repository <local repository path>
Delete from the list of local repository paths:	
	gitworklog config delete-repository <local repository path>
List registered repository paths
	gitworklog config list-repositories
display registered git user name
	gitworklog config get-name"`,
	)
}
