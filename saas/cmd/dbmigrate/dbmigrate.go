package main

import (
	"configmgr"
	"context"
	"log"
	"saas/gen/ent"
	"saas/gen/ent/migrate"
	"saas/pkg/app"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

func main() {
	cfg := configmgr.GetAppConfig()
	db := app.InitDB(cfg)

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv), ent.Debug())

	defer client.Close()
	// Run the auto migration tool.
	err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false), // Disable foreign keys.
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
