package cli

import (
	"strings"

	"github.com/spf13/cobra"
)

type App struct {
	ctx *Context
	cmd *cobra.Command
	subApp *App
	isRootApp bool
}

func New(initConfig func()) *App {
	if initConfig != nil {
		cobra.OnInitialize(initConfig)
	}

	ctx := &Context{}

	return &App{ctx: ctx, cmd: &cobra.Command{}, isRootApp: true}
}

func (a *App) AddCommand(cmdString string, run CommandRunFuncWithCtx, meta Meta, initFunc Init) *App {
	cmdString = strings.TrimSpace(cmdString)
	
	subCmd := &cobra.Command{
		Use: cmdString,
		Run: func(cmd *cobra.Command, args []string) {
			run(a.ctx, cmd, args)
		},
	}

	subCmd.Short = meta.Short
	subCmd.Long = meta.Long

	if initFunc != nil {
		initFunc(subCmd)
	}
	
	a.cmd.AddCommand(subCmd)

	a.subApp = &App{ctx: a.ctx, cmd: subCmd}

	return a.subApp
}

func (a *App) ApplyPreRun(options ...Option) *App {
	cmd := a.cmd
	if a.isRootApp {
		cmd = a.subApp.cmd
	}

	currentPreRun := cmd.PreRun

	for _, option := range options {
		currentPreRun = option(currentPreRun)
	}

	cmd.PreRun = currentPreRun

	return a
}

func (a *App) ApplyPostRun(options ...Option) *App {
	cmd := a.cmd
	if a.isRootApp {
		cmd = a.subApp.cmd
	}

	currentPostRun := cmd.PostRun

	for _, option := range options {
		currentPostRun = option(currentPostRun)
	}

	cmd.PostRun = currentPostRun

	return a
}

func (a *App) ApplyPreRunE(options ...OptionE) *App {
	cmd := a.cmd
	if a.isRootApp {
		cmd = a.subApp.cmd
	}

	currentPreRunE := cmd.PreRunE

	for _, option := range options {
		currentPreRunE = option(currentPreRunE)
	}

	cmd.PreRunE = currentPreRunE

	return a
}

func (a *App) ApplyPostRunE(options ...OptionE) *App {
	cmd := a.cmd
	if a.isRootApp {
		cmd = a.subApp.cmd
	}

	currentPostRunE := cmd.PreRunE

	for _, option := range options {
		currentPostRunE = option(currentPostRunE)
	}

	cmd.PostRunE = currentPostRunE

	return a
}

func (a *App) Execute() error {
	err := a.cmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
