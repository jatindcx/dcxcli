package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"dcxcli/internal/generate"
	"dcxcli/internal/tgscip"
	"dcxcli/pkg/cli"
)

var rootCmd = &cobra.Command{}

// func main() {
// 	// rootCmd.AddCommand(generate.CmdGenerate)
// 	tgscip.CmdTgscip.PreRun = func(cmd *cobra.Command, args []string) {
// 		fmt.Println("Args: ", args)
// 	}

// 	rootCmd.AddCommand(tgscip.CmdTgscip)
// 	// tgscip.CmdTgscip.AddCommand(generate.CmdGenerate)

// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Println(err)
// 	}
// }

func main() {
	app := cli.New(nil)

	app.AddCommand("tgscip", tgscip.DoSomething, cli.Meta{}, nil).AddCommand("generate", generate.GeneratePassword, cli.Meta{}, generate.Init)
	app.ApplyPreRun(SamplePreRun, SamplePreRun)

	if err := app.Execute(); err != nil {
		fmt.Println(err)
	}
}

func SamplePreRun(next cli.CommandRunFunc) cli.CommandRunFunc {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("SamplePreRun: ", args)

		if next != nil {
			next(cmd, args)
		}
	}
}