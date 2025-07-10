package tgscip

import (
	"github.com/spf13/cobra"

	"dcxcli/pkg/cli"
)

func DoSomething(ctx *cli.Context, cmd *cobra.Command, args []string) {
	ctx.Logger.Info("Hey there!, you are at TGSCIP")
}