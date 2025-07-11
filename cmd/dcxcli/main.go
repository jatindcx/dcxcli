package main

import (
	"fmt"
	"dcxcli/pkg/cli"
)

func main() {
	app := cli.New(nil)
	InitService(app)
	if err := app.Execute(); err != nil {
		fmt.Println(err)
	}
}
