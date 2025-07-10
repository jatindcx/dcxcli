package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"dcxcli/internal/generate"
	"dcxcli/internal/tgscip"
)

var rootCmd = &cobra.Command{}

func main() {
	rootCmd.AddCommand(generate.CmdGenerate)
	rootCmd.AddCommand(tgscip.CmdTgscip)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}