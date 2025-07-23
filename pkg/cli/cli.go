package cli

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"dcxcli/pkg/preRun"
	"dcxcli/pkg/types"
)

type App struct {
	ctx       *types.Context
	cmd       *cobra.Command
	subApp    *App
	isRootApp bool
}

// New factory function to create a new App instance.
func New(initConfig func()) *App {
	if initConfig != nil {
		cobra.OnInitialize(initConfig)
	}

	logger, _ := zap.NewDevelopment()
	ctx := &types.Context{Logger: logger}

	return &App{ctx: ctx, cmd: &cobra.Command{}, isRootApp: true}
}

// AddCommand registers a new command. Arguments required are name of command, command function, Meta (Long description of command),
// initializer function. User can write initializer function to register flags.
func (a *App) AddCommand(cmdString string, run types.CommandRunFuncWithCtx, meta types.Meta, initFunc types.Init) *App {
	cmdString = strings.TrimSpace(cmdString)

	subCmd := &cobra.Command{
		Use: cmdString,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := a.ctx
			ctx.Ctx = cmd.Context()

			run(ctx, cmd, args)
		},
	}

	subCmd.Long = meta.Long

	if initFunc != nil {
		initFunc(subCmd)
	}

	a.cmd.AddCommand(subCmd)

	a.subApp = &App{ctx: a.ctx, cmd: subCmd}

	a.subApp.ApplyPreRun(preRun.Auth)

	return a.subApp
}

// ApplyPreRun adds PreRun function(s).
func (a *App) ApplyPreRun(options ...types.OptionWithCtx) *App {
	app := a
	if a.isRootApp {
		app = a.subApp
	}

	cmd := app.cmd

	currentPreRun := cmd.PreRun

	for _, option := range options {
		tempPreRun := currentPreRun
		currentPreRun = func(cmd *cobra.Command, args []string) {
			ctx := app.ctx
			ctx.Ctx = context.Background()

			option(tempPreRun)(ctx, cmd, args)
		}
	}

	cmd.PreRun = currentPreRun

	return a
}

// ApplyPostRun add PostRun function(s).
func (a *App) ApplyPostRun(options ...types.OptionWithCtx) *App {
	app := a
	if a.isRootApp {
		app = a.subApp
	}

	cmd := app.cmd

	currentPostRun := cmd.PostRun

	for _, option := range options {
		tempPostRun := currentPostRun
		currentPostRun = func(cmd *cobra.Command, args []string) {
			ctx := app.ctx
			ctx.Ctx = context.Background()

			option(tempPostRun)(ctx, cmd, args)
		}
	}

	cmd.PostRun = currentPostRun

	return a
}

// ApplyPreRunE adds PreRunE functions(s).
func (a *App) ApplyPreRunE(options ...types.OptionEWithCtx) *App {
	app := a
	if a.isRootApp {
		app = a.subApp
	}

	cmd := app.cmd

	currentPreRunE := cmd.PreRunE

	for _, option := range options {
		tempPreRunE := currentPreRunE
		currentPreRunE = func(cmd *cobra.Command, args []string) error {
			ctx := app.ctx
			ctx.Ctx = context.Background()

			return option(tempPreRunE)(ctx, cmd, args)
		}
	}

	cmd.PreRunE = currentPreRunE

	return a
}

// ApplyPostRunE adds PostRunE function(s).
func (a *App) ApplyPostRunE(options ...types.OptionEWithCtx) *App {
	app := a
	if a.isRootApp {
		app = a.subApp
	}

	cmd := app.cmd

	currentPostRunE := cmd.PreRunE

	for _, option := range options {
		tempPostRunE := currentPostRunE
		currentPostRunE = func(cmd *cobra.Command, args []string) error {
			ctx := app.ctx
			ctx.Ctx = context.Background()

			return option(tempPostRunE)(ctx, cmd, args)
		}
	}

	cmd.PostRunE = currentPostRunE

	return a
}

// Execute triggers cobra.Command.Execute().
func (a *App) Execute() error {
	err := a.cmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
