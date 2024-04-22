package resolverfn

import (
	"configmgr/gen/appconfig"
	"saas/pkg/app"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Plugin    app.Plugin
	AppConfig *appconfig.AppConfig
}
