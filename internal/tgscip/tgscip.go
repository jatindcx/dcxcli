package tgscip

import (
	"fmt"

	"github.com/spf13/cobra"

	"dcxcli/pkg/cli"
)

func DoSomething(ctx *cli.Context, cmd *cobra.Command, args []string) {
	fmt.Println("Hey there!, you are at TGSCIP")
}