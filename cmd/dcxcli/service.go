package main

import (

	"dcxcli/internal/generate"
	"dcxcli/internal/tgscip"
	"dcxcli/pkg/cli"
	"dcxcli/pkg/types"
	"dcxcli/internal/mock"

)

func InitService(app *cli.App) {

	// adding nested command, e.g. dcxcli tgscip generate
	// AddCommand needs, use, run, Meta(Short and Long), Init(for initializing flags

	app.AddCommand("tgscip", tgscip.DoSomething, types.Meta{}, nil).AddCommand("generate", generate.GeneratePassword, types.Meta{}, generate.Init)
	
	app.AddCommand(
		"mock",
		mock.DoSomething,
		types.Meta{Long: "Simulate Docker image pull with mock"},
		mock.Init,
	)	
}