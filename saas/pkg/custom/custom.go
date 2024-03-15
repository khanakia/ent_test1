package custom

import (
	"fmt"
	"lace/nlog"

	"lace/cli"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// This package is more like global and we will register this package only as plugin
// All the GormDB Auto Migration table should go there

type Config struct {
	// GormDB gormdb.GormDB
	Logger nlog.Logger
	Cli    cli.Cli
	// Server server.Server
}

type Acore struct {
	Config
}

func (pkg Acore) Version() string {
	return "0.01"
}

func New(config Config) Acore {
	pkg := Acore{
		Config: config,
	}
	AddCommands(config)

	return pkg
}

func AddCommands(config Config) {
	var Main = &cobra.Command{
		Use:   "core",
		Short: "Core Pkg - Use `go run . ana --help` to see child commands",
		Run: func(cmd *cobra.Command, args []string) {
			color.Yellow("Run below command to see all the child commands.")
			color.Cyan("go run . core --help")
		},
	}

	rootCmd := config.Cli.RootCmd
	rootCmd.AddCommand(Main)

	Main.AddCommand(&cobra.Command{
		Use:   "test",
		Short: "Test",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("All Done!!!")

		},
	})

}
