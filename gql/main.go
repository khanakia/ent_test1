package gql

import (
	"fmt"
	"gql/graph/generated"
	"gql/graph/resolverfn"
	"gql/internal/gqlgin"
	"lace/gqlgenfn"
	"saas/pkg/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func Boot(ginEngine *gin.Engine, resolver *resolverfn.Resolver) {
	prefix := "/app"
	rg := ginEngine.Group(prefix)

	// ginEngine.Use(middleware.MiddlewareSilent(resolver.Plugin.EntDB.Client()))
	// ginEngine.Use(gqlgenfn.GinContextToContextMiddleware())

	c := generated.Config{Resolvers: resolver}

	gserver := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	gqlgin.New(gqlgin.Config{
		RouterGroup:      rg,
		GqlServer:        gserver,
		PlaygroundKey:    resolver.AppConfig.Graphql.Key,
		RouteGroupPrefix: prefix,
		Middleware: []gin.HandlerFunc{
			middleware.MiddlewareSilent(resolver.Plugin.EntDB.Client()),
			gqlgenfn.GinContextToContextMiddleware(),
		},
	})
	fmt.Println("boot gql")
}
