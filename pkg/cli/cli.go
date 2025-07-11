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

func New(initConfig func()) *App {
	if initConfig != nil {
		cobra.OnInitialize(initConfig)
	}

	logger, _ := zap.NewDevelopment()
	ctx := &types.Context{Logger: logger}

	return &App{ctx: ctx, cmd: &cobra.Command{}, isRootApp: true}
}

func (a *App) AddCommand(cmdString string, run types.CommandRunFuncWithCtx, meta types.Meta, initFunc types.Init) *App {
	cmdString = strings.TrimSpace(cmdString)

	subCmd := &cobra.Command{
		Use: cmdString,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := a.ctx
			ctx.Ctx = context.Background()

			run(ctx, cmd, args)
		},
	}

	subCmd.Short = meta.Short
	subCmd.Long = meta.Long

	if initFunc != nil {
		initFunc(subCmd)
	}

	a.cmd.AddCommand(subCmd)

	a.subApp = &App{ctx: a.ctx, cmd: subCmd}

	a.subApp.ApplyPreRun(preRun.Auth)

	return a.subApp
}

func (a *App) ApplyPreRun(options ...types.OptionWithCtx) *App {
	cmd := a.cmd
	if a.isRootApp {
		cmd = a.subApp.cmd
	}

	currentPreRun := cmd.PreRun

	for _, option := range options {
		tempPreRun := currentPreRun
		currentPreRun = func(cmd *cobra.Command, args []string) {
			ctx := a.ctx
			ctx.Ctx = context.Background()

			option(tempPreRun)(ctx, cmd, args)
		}
	}

	cmd.PreRun = currentPreRun

	return a
}

func (a *App) ApplyPostRun(options ...types.OptionWithCtx) *App {
	cmd := a.cmd
	if a.isRootApp {
		cmd = a.subApp.cmd
	}

	currentPostRun := cmd.PostRun

	for _, option := range options {
		tempPostRun := currentPostRun
		currentPostRun = func(cmd *cobra.Command, args []string) {
			ctx := a.ctx
			ctx.Ctx = context.Background()

			option(tempPostRun)(ctx, cmd, args)
		}
	}

	cmd.PostRun = currentPostRun

	return a
}

func (a *App) ApplyPreRunE(options ...types.OptionEWithCtx) *App {
	cmd := a.cmd
	if a.isRootApp {
		cmd = a.subApp.cmd
	}

	currentPreRunE := cmd.PreRunE

	for _, option := range options {
		tempPreRunE := currentPreRunE
		currentPreRunE = func(cmd *cobra.Command, args []string) error {
			ctx := a.ctx
			ctx.Ctx = context.Background()

			return option(tempPreRunE)(ctx, cmd, args)
		}
	}

	cmd.PreRunE = currentPreRunE

	return a
}

func (a *App) ApplyPostRunE(options ...types.OptionEWithCtx) *App {
	cmd := a.cmd
	if a.isRootApp {
		cmd = a.subApp.cmd
	}

	currentPostRunE := cmd.PreRunE

	for _, option := range options {
		tempPostRunE := currentPostRunE
		currentPostRunE = func(cmd *cobra.Command, args []string) error {
			ctx := a.ctx
			ctx.Ctx = context.Background()

			return option(tempPostRunE)(ctx, cmd, args)
		}
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
