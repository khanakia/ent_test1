package gqlsa

import (
	"fmt"
	"gqlsa/graph/generated"
	"gqlsa/graph/resolverfn"
	"gqlsa/internal/gqlgin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func Boot(ginEngine *gin.Engine) {

	resolver := &resolverfn.Resolver{
		// GormDB:    plugin.GormDB,
		// Logger:    plugin.Logger,
		// Cache:     plugin.CacheRdbms,
		// EntClient: plugin.EntDB.Client(),
	}

	c := generated.Config{Resolvers: resolver}

	gserver := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	// gserver := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: config.Resolver}))

	gqlgin.New(gqlgin.Config{
		GinEngine: ginEngine,
		GqlServer: gserver,
		// Middleware: []gin.HandlerFunc{middleware.JwtMiddlewarePublic(config.GormDB.DB)},
	})
	fmt.Println("boot")
}
