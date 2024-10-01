package main

import (
	"configmgr"
	"configmgr/gen/appconfig/environment"
	"gql"
	"gql/graph/resolverfn"
	"gqlsa"
	"gqlsa/graph/gqlsaresolver"
	"lace/ginserver"
	"lace/natso"
	"lace/util"
	"saas/cmd/api/natshandlers"
	"saas/pkg/app"
	"saas/pkg/apphandlers"
)

func main() {
	appcfg := configmgr.GetAppConfig()

	serverServer := ginserver.New(
		ginserver.WithBeforeHandler(util.CORSMiddleware()),
		ginserver.WithAppName(appcfg.Hostname),
		ginserver.WithPort(appcfg.Port),
		ginserver.WithIsProd(appcfg.Environment == environment.Prod),
	)

	plugins := app.New()

	// add this line to init or add the graphql to the api
	gql.Boot(serverServer.Router, &resolverfn.Resolver{
		Plugin:    app.GetPlugins(),
		AppConfig: appcfg,
	})

	gqlsa.Boot(serverServer.Router, &gqlsaresolver.Resolver{
		Plugin:    app.GetPlugins(),
		AppConfig: appcfg,
	})

	if appcfg.Nats.Status {
		natsnats := natso.New(natso.WithAppKey(appcfg.Nats.AppKey), natso.WithURL(appcfg.Nats.Url), natso.WithEnabled(true))
		natshandlers.New(natshandlers.Config{
			Natso: natsnats,
			// EntClient: appApp.EntDB.Client(),
			// Nlog:      appApp.Nlog,
		})
	}

	apphandlers.New(apphandlers.Config{
		Server:    serverServer,
		EntClient: plugins.EntDB.Client(),
		Nlog:      plugins.Nlog,
		AppConfig: plugins.AppConfig,
	})

	serverServer.Start()
}
