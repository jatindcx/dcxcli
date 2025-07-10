package preRun

import (
	"fmt"

	"github.com/spf13/cobra"

	"dcxcli/pkg/cli"
)

func SamplePreRun(next cli.CommandRunFunc) cli.CommandRunFunc {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("SamplePreRun: ", args)

		if next != nil {
			next(cmd, args)
		}
	}
}