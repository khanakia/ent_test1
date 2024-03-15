package main

import "saas/pkg/app"

func main() {
	app := app.New()
	app.Cli.Execute()
}
