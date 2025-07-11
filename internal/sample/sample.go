package sample

import (
	"dcxcli/pkg/types"
	"fmt"
	"github.com/spf13/cobra"
)

func Init(sampleCmd *cobra.Command) {
	// define flags
	sampleCmd.Flags().IntP("length", "l", 0, "Length of password")
}

func SampleCommand(ctx *types.Context, cmd *cobra.Command, args []string) {
	// create command
	ctx.Logger.Info("Sample Command Started")

	length, _ := cmd.Flags().GetInt("length")

	fmt.Println("Length = ", length)

	//ctx.Logger.Info("Arguments Provided: ", zap.Field{Key: "l", })
	//time.Sleep(2 * time.Second)

	ctx.Logger.Info("Sample Command Completed")
}

func SampleNestedCommand(ctx *types.Context, cmd *cobra.Command, args []string) {
	ctx.Logger.Info("Sample Nested Command Started")

	ctx.Logger.Info("Sample Nested Command Completed")
}

func SamplePreRun(currPrerun types.CommandRunFunc) types.CommandRunFuncWithCtx {
	return func(ctx *types.Context, cmd *cobra.Command, args []string) {
		ctx.Logger.Info("Sample PreRun Started")
	}
}

func SamplePreRunE(currPrerun types.CommandRunEFunc) types.CommandRunEFuncWithCtx {
	return func(ctx *types.Context, cmd *cobra.Command, args []string) error {
		ctx.Logger.Info("Sample PreRun Started")

		return fmt.Errorf("Sample PreRun Failed")
	}
}
