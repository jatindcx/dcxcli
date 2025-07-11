package types

import "github.com/spf13/cobra"

type (
	CommandRunFunc  func(*cobra.Command, []string)
	CommandRunEFunc func(*cobra.Command, []string) error

	CommandRunFuncWithCtx  func(*Context, *cobra.Command, []string)
	CommandRunEFuncWithCtx func(*Context, *cobra.Command, []string) error

	Option  func(CommandRunFunc) CommandRunFunc
	OptionE func(CommandRunEFunc) CommandRunEFunc

	OptionWithCtx  func(CommandRunFunc) CommandRunFuncWithCtx
	OptionEWithCtx func(CommandRunEFunc) CommandRunEFuncWithCtx

	Init func(*cobra.Command)
)
