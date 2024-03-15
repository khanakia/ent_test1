package app

import (
	"configmgr"
	"configmgr/gen/appconfig"
	"database/sql"
	"fmt"
	"saas/pkg/entdb"

	"lace/db"
	"lace/nlog"

	"lace/cli"
)

func New() Plugin {
	cli := cli.New()

	appcfg := configmgr.GetAppConfig()

	db := InitDB(appcfg)

	return Plugin{
		Cli:   cli,
		Nlog:  nlog.New(),
		DB:    db,
		EntDB: entdb.New(db),
	}
}

type Plugin struct {
	Cli   cli.Cli
	Nlog  nlog.Logger
	DB    db.DB
	EntDB entdb.EntDB
}

func InitDB(cfg *appconfig.AppConfig) *sql.DB {
	return db.New(db.WithDataSourceName(dsName(cfg)), db.WithDriverName("pgx"))
}

func dsName(cfg *appconfig.AppConfig) string {
	username := cfg.Database.Username
	password := cfg.Database.Password
	databaseName := cfg.Database.DbName
	databaseHost := cfg.Database.Host
	databasePort := cfg.Database.Port
	sslmode := cfg.Database.Sslmode

	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s port=%d", databaseHost, username, databaseName, sslmode, password, databasePort)
}
