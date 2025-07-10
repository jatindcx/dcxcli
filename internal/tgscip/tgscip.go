package tgscip

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdTgscip *cobra.Command

func init() {
	CmdTgscip = &cobra.Command{
		Use: "tgscip",
		Short: "Sample Help",
		Long: "Sample Long Help",
		Run: doSomething,
	}
}

func doSomething(cmd *cobra.Command, args []string) {
	fmt.Println("Hey there!, you are at TGSCIP")
}