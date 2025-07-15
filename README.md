# DCXCLI

A Go CLI framework built on top of Cobra that provides a simplified interface for creating command-line applications with built-in context management and middleware support.

## Features

- Simple command registration
- Built-in context management with logging
- PreRun and PostRun middleware support
- Error handling with PreRunE and PostRunE

## Quick Start

```go
package main

import (
    "dcxcli/pkg/cli"
    "fmt"
)

func main() {
    app := cli.New(nil)
    
    // Register your commands
    InitService(app)
    
    if err := app.Execute(); err != nil {
        fmt.Println(err)
    }
}
```

## Registering Commands

Use the `AddCommand` method to register new commands:

```go
func InitService(app *cli.App) {
    app.AddCommand(
        "mock",                    // command name
        mock.MockCommand,          // command handler function
        types.Meta{Long: "Simulate Docker image pull with mock"}, // metadata
        mock.Init,                 // initialization function (optional)
    )
}
```

### Command Handler Function

Your command handler must implement `types.CommandRunFuncWithCtx`:

```go
func MockCommand(ctx *types.Context, cmd *cobra.Command, args []string) {
    ctx.Logger.Info("Command executed")
    // Your command logic here
}
```

### Initialization Function

The init function allows you to configure flags and other command properties:

```go
func Init(cmd *cobra.Command) {
    cmd.Flags().StringVarP(&imageName, "image", "i", "", "Docker image name")
    cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed logs")
}
```

## PreRun Functions

PreRun functions execute before the main command handler. They're useful for authentication, validation, or setup tasks.

### Applying PreRun Functions

```go
// Apply a single PreRun function
app.AddCommand("mycommand", handler, meta, init).
    ApplyPreRun(preRun.Auth)

// Apply multiple PreRun functions (they execute in order)
app.AddCommand("mycommand", handler, meta, init).
    ApplyPreRun(preRun.Auth, preRun.Validate)
```

### Creating PreRun Functions

PreRun functions must implement `types.OptionWithCtx`:

```go
func Auth(next types.CommandRunFunc) types.CommandRunFuncWithCtx {
    return func(ctx *types.Context, cmd *cobra.Command, args []string) {
        ctx.Logger.Info("Auth PreRun")
        
        // Your authentication logic here
        if !isAuthenticated() {
            ctx.Logger.Error("Authentication failed")
            return
        }
        
        // Call the next function in the chain
        if next != nil {
            next(cmd, args)
        }
    }
}
```

### PreRunE (with Error Handling)

For PreRun functions that can return errors:

```go
app.AddCommand("mycommand", handler, meta, init).
    ApplyPreRunE(preRun.AuthWithError)

func AuthWithError(next types.CommandRunEFunc) types.CommandRunEFuncWithCtx {
    return func(ctx *types.Context, cmd *cobra.Command, args []string) error {
        if !isAuthenticated() {
            return fmt.Errorf("authentication failed")
        }
        
        if next != nil {
            return next(cmd, args)
        }
        return nil
    }
}
```

## PostRun Functions

PostRun functions execute after the main command handler, useful for cleanup or logging.

### Applying PostRun Functions

```go
app.AddCommand("mycommand", handler, meta, init).
    ApplyPostRun(postRun.Cleanup, postRun.LogStats)
```

### Creating PostRun Functions

PostRun functions follow the same pattern as PreRun:

```go
func Cleanup(next types.CommandRunFunc) types.CommandRunFuncWithCtx {
    return func(ctx *types.Context, cmd *cobra.Command, args []string) {
        // Call the next function first (if any)
        if next != nil {
            next(cmd, args)
        }
        
        // Your cleanup logic here
        ctx.Logger.Info("Cleanup completed")
    }
}
```

### PostRunE (with Error Handling)

```go
app.AddCommand("mycommand", handler, meta, init).
    ApplyPostRunE(postRun.CleanupWithError)
```

## Chaining Middleware

You can chain multiple PreRun and PostRun functions:

```go
app.AddCommand("mycommand", handler, meta, init).
    ApplyPreRun(preRun.Auth, preRun.Validate).
    ApplyPostRun(postRun.Cleanup, postRun.LogStats)
```

## Context

The framework provides a context object with built-in logging:

```go
func MyCommand(ctx *types.Context, cmd *cobra.Command, args []string) {
    ctx.Logger.Info("Starting command")
    ctx.Logger.Error("Something went wrong", zap.Error(err))
    
    // Access the underlying context.Context
    select {
    case <-ctx.Ctx.Done():
        ctx.Logger.Info("Command cancelled")
        return
    default:
        // Continue processing
    }
}
```

## Example: Complete Command with Middleware

```go
// Command handler
func ProcessData(ctx *types.Context, cmd *cobra.Command, args []string) {
    ctx.Logger.Info("Processing data...")
    // Your processing logic here
}

// Initialization
func InitProcessCommand(cmd *cobra.Command) {
    cmd.Flags().StringP("input", "i", "", "Input file path")
    cmd.Flags().StringP("output", "o", "", "Output file path")
}

// Registration with middleware
func InitService(app *cli.App) {
    app.AddCommand(
        "process",
        ProcessData,
        types.Meta{Long: "Process data from input to output"},
        InitProcessCommand,
    ).ApplyPreRun(preRun.Auth, preRun.ValidateFiles).
      ApplyPostRun(postRun.Cleanup)
}
```

## Building and Running

```bash
# Build the CLI
go build -o dcxcli ./cmd/dcxcli

# Run a command
./dcxcli mock --image nginx:latest --verbose
```