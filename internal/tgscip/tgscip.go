package tgscip

import (
	"github.com/spf13/cobra"

	"dcxcli/pkg/types"
)

func DoSomething(ctx *types.Context, cmd *cobra.Command, args []string) {
	ctx.Logger.Info("Hey there!, you are at TGSCIP")
}
