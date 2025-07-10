package main

import (
	"fmt"

	"dcxcli/internal/generate"
	"dcxcli/internal/tgscip"
	"dcxcli/pkg/cli"
	"dcxcli/pkg/preRun"
)

func main() {
	app := cli.New(nil)

	// adding nested command, e.g. dcxcli tgscip generate
	// AddCommand needs, use, run, Meta(Short and Long), Init(for initializing flags)
	app.AddCommand("tgscip", tgscip.DoSomething, cli.Meta{}, nil).AddCommand("generate", generate.GeneratePassword, cli.Meta{}, generate.Init)
	app.ApplyPreRun(preRun.SamplePreRun)

	if err := app.Execute(); err != nil {
		fmt.Println(err)
	}
}