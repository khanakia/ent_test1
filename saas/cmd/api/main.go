package main

import (
	"configmgr"
	"configmgr/gen/appconfig/environment"
	"gqlsa"
	"lace/ginserver"
	"lace/util"
)

func main() {

	appcfg := configmgr.GetAppConfig()

	serverServer := ginserver.New(
		ginserver.WithBeforeHandler(util.CORSMiddleware()),
		ginserver.WithAppName(appcfg.Hostname),
		ginserver.WithPort(appcfg.Port),
		ginserver.WithIsProd(appcfg.Environment == environment.Prod),
	)

	gqlsa.Boot(serverServer.Router)

	serverServer.Start()
}
