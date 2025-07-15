package main

import (
	"dcxcli/internal/mock"
	"dcxcli/pkg/cli"
	"dcxcli/pkg/types"
)

func InitService(app *cli.App) {
	// registering command "mock"
	app.AddCommand(
		"mock",
		mock.MockCommand,
		types.Meta{Long: "Simulate Docker image pull with mock"},
		mock.Init,
	)
}
