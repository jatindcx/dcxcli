package preRun

import (
	"github.com/spf13/cobra"

	"dcxcli/pkg/types"
)

func Auth(next types.CommandRunFunc) types.CommandRunFuncWithCtx {
	return func(ctx *types.Context, cmd *cobra.Command, args []string) {
		ctx.Logger.Info("Auth PreRun")

		if next != nil {
			next(cmd, args)
		}
	}
}
