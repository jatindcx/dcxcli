package main

import (
	"dcxcli/internal/sample"
	"dcxcli/pkg/types"
	"fmt"

	"dcxcli/pkg/cli"
)

// 1. Import the new repo
// 2. Create handler in internal

func main() {
	app := cli.New(nil)

	app.AddCommand("sample", sample.SampleCommand, types.Meta{Short: "Short desc", Long: "Long Desc"}, sample.Init).
		AddCommand("nested", sample.SampleNestedCommand, types.Meta{Short: "Short nested desc", Long: "Long nested Desc"}, nil)

	// apply pre run
	//app.ApplyPreRun(sample.SamplePreRun)

	// apply pre run with error
	app.ApplyPreRunE(sample.SamplePreRunE)

	if err := app.Execute(); err != nil {
		fmt.Println(err)
	}
}
