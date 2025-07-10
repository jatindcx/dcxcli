package cli

import (
	"strings"

	"github.com/spf13/cobra"
)

type App struct {
	cmd *cobra.Command
	subApp *App
}

func New(initConfig func()) *App {
	if initConfig != nil {
		cobra.OnInitialize(initConfig)
	}

	return &App{cmd: &cobra.Command{}}
}

func (a *App) AddCommand(cmdString string, run CommandRunFunc, meta Meta, initFunc Init) *App {
	cmdString = strings.TrimSpace(cmdString)
	
	subCmd := &cobra.Command{
		Use: cmdString,
		Run: run,
	}

	subCmd.Short = meta.Short
	subCmd.Long = meta.Long

	if initFunc != nil {
		initFunc(subCmd)
	}
	
	a.cmd.AddCommand(subCmd)
	a.subApp = &App{cmd: subCmd}

	return a.subApp
}

func (a *App) ApplyPreRun(options ...Option) *App {
	cmd := a.cmd
	if len(cmd.Use) == 0 {
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
	if len(cmd.Use) == 0 {
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
	if len(cmd.Use) == 0 {
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
	if len(cmd.Use) == 0 {
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
