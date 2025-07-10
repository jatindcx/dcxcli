package cli

import (
	"github.com/spf13/cobra"
)

type (
	CommandRunFunc func(*cobra.Command, []string)
	CommandRunEFunc func(*cobra.Command, []string) error
	Option func(CommandRunFunc) CommandRunFunc
	OptionE func(CommandRunEFunc) CommandRunEFunc
	Init func(*cobra.Command)
)

type Meta struct {
	Short string
	Long string
}