package apphandlers

import (
	"configmgr/gen/appconfig"
	"lace/ginserver"
	"lace/nlog"
	"saas/gen/ent"
)

type Config struct {
	Nlog      nlog.Logger
	Server    ginserver.Server
	EntClient *ent.Client
	AppConfig *appconfig.AppConfig
}

type AppHandlers struct {
	Config
}

func New(config Config) {
	// router := config.Server.Router
}
