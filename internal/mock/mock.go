package mock

import (
	"fmt"

	"github.com/jatindcx/dcxcli-mock-test/mock"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"dcxcli/pkg/types"
)

var imageName string
var verbose bool

// MockCommand is a mock command that executes the function from mock repo.
func MockCommand(ctx *types.Context, _ *cobra.Command, _ []string) {
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
