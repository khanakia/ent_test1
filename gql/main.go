package gql

import (
	"fmt"
	"gql/graph/generated"
	"gql/graph/resolverfn"
	"lace/gqlgin"
	"saas/pkg/middleware"
	"saas/pkg/middleware/appmiddleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func Boot(ginEngine *gin.Engine, resolver *resolverfn.Resolver) {
	prefix := "/app"
	rg := ginEngine.Group(prefix)

	c := generated.Config{Resolvers: resolver}
	// setDirectives(&c)

	schm := generated.NewExecutableSchema(c)
	// extendSchema(&schm)
	gserver := handler.NewDefaultServer(schm)

	gqlgin.New(gqlgin.Config{
		RouterGroup:      rg,
		GqlServer:        gserver,
		PlaygroundKey:    resolver.AppConfig.Graphql.Key,
		RouteGroupPrefix: prefix,
		Middleware: []gin.HandlerFunc{
			appmiddleware.CheckAppMiddleware(resolver.Plugin.EntDB.Client()),
			middleware.MiddlewareSilent(resolver.Plugin.EntDB.Client()),
		},
	})
	fmt.Println("boot gql")
}
