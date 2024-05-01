package main

import (
	"saas/pkg/app"
)

func main() {
	app := app.New()
	app.Cli.Execute()
	// fmt.Println(authfn.Hash("admin"))
}
