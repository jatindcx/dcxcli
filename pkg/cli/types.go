package cli

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

type (
	CommandRunFunc func(*cobra.Command, []string)
	CommandRunEFunc func(*cobra.Command, []string) error

	CommandRunFuncWithCtx func(*Context, *cobra.Command, []string)
	CommandRunEFuncWithCtx func(*Context, *cobra.Command, []string) error

	Option func(CommandRunFunc) CommandRunFunc
	OptionE func(CommandRunEFunc) CommandRunEFunc

	Init func(*cobra.Command)
)

type Meta struct {
	Short string
	Long string
}

type Context struct {
	ctx context.Context
	Logger log.Logger
}