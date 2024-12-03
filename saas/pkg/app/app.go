package app

import (
	"configmgr"
	"configmgr/gen/appconfig"
	"database/sql"
	"fmt"
	"saas/pkg/cacherdbms"
	"saas/pkg/entdb"

	"lace/cache"
	"lace/db"
	"lace/medialibrary"
	"lace/nlog"

	"lace/cli"

	_ "saas/gen/ent/runtime"
)

var plugins Plugin

func New() Plugin {
	cli := cli.New()

	appcfg := configmgr.GetAppConfig()

	db := InitDB(appcfg)

	entClient := entdb.New(db)

	plugins = Plugin{
		Cli:       cli,
		Nlog:      nlog.New(appcfg.LogFile),
		DB:        db,
		EntDB:     entClient,
		AppConfig: appcfg,
		Cache:     cacherdbms.New(entClient.Client()),
	}

	medialibrary.Construct(db)

	return plugins
}

func GetPlugins() Plugin {
	return plugins
}

type Plugin struct {
	Cli       cli.Cli
	Nlog      nlog.Logger
	DB        db.DB
	EntDB     entdb.EntDB
	Cache     cache.Store
	AppConfig *appconfig.AppConfig
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
