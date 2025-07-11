package preRun

import (
	"fmt"

	"github.com/spf13/cobra"

	"dcxcli/pkg/cli"
)

func SamplePreRun(next cli.CommandRunFunc) cli.CommandRunFuncWithCtx {
	return func(ctx *cli.Context, cmd *cobra.Command, args []string) {
		fmt.Println("SamplePreRun: ", args)

		if next != nil {
			next(cmd, args)
		}
	}
}
