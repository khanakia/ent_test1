package main

import (
	"saas/pkg/app"
	"saas/pkg/cacherdbms"
)

func main() {
	app := app.New()
	app.Cli.Execute()
	cacherd := cacherdbms.New(app.EntDB.Client())
	cacherd.Put("name", "luci", 1000)
}
