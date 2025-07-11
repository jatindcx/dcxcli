package mock

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/jatindcx/dcxcli-mock-test/mock"
	"dcxcli/pkg/types"
	"go.uber.org/zap"
)

var imageName string
var verbose bool

func DoSomething(ctx *types.Context, cmd *cobra.Command, args []string) {

	imageName, _ := cmd.Flags().GetString("image")
	verbose, _ := cmd.Flags().GetBool("verbose")

	opts := mock.MockOptions{
		Image:   imageName,
		Verbose: verbose,
	}
	
	ctx.Logger.Info("Pulling image", zap.String("image", opts.Image))

	result := mock.PullImage(opts, ctx.Logger)

	if result.Err != nil {
		ctx.Logger.Error("Error:", zap.Error(result.Err))
		fmt.Printf("Exit Code: %d\n", result.StatusCode)
	} else {
		ctx.Logger.Info("Success:", zap.String("output", result.Output))
	}
}

func Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&imageName, "image", "i", "", "Docker image name to simulate pull")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed pull logs")
}
