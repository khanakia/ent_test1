package gqlgin

import (
	"context"
	"fmt"

	"lace/gqlgenfn"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type Gql struct {
	Config
}

const PlaygroundKey = "pkey"

type Config struct {
	RouterGroup      *gin.RouterGroup
	GqlServer        *handler.Server
	Middleware       []gin.HandlerFunc
	RouteGroupPrefix string
	PlaygroundKey    string
}

// Defining the Graphql handler
func (cls *Gql) instrospectionMiddleware(gserver *handler.Server) gin.HandlerFunc {
	// https://github.com/99designs/gqlgen/blob/master/docs/content/reference/introspection.md
	gserver.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		// isAllowedPlayground := gqlgenfn.IsPlaygroundAllowedOperation(cls.PlaygroundKey, ctx)
		isAllowedPlayground := gqlgenfn.IsPlaygroundGqlCtx(ctx)
		fmt.Println("isAllowedPlayground", isAllowedPlayground)
		if isAllowedPlayground {
			graphql.GetOperationContext(ctx).DisableIntrospection = false
		}

		return next(ctx)
	})

	return func(c *gin.Context) {
		c.Next()
	}
}

func (cls *Gql) serverHttpHandler(gserver *handler.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		gserver.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler(routeGroupPrefix string) gin.HandlerFunc {
	h := playground.Handler("GraphQL", fmt.Sprintf("%s/query_playground", routeGroupPrefix))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func New(config Config) Gql {
	gqlgin := Gql{
		Config: config,
	}

	routeGroupPrefix := ""
	if len(config.RouteGroupPrefix) > 0 {
		routeGroupPrefix = config.RouteGroupPrefix
	}

	rg := config.RouterGroup

	// disable the introspection by default
	config.GqlServer.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		graphql.GetOperationContext(ctx).DisableIntrospection = true
		return next(ctx)
	})

	// middleware orders is very important so do not change it
	// for making actual request and can have userAuth and other Middleware in between
	config.Middleware = append(config.Middleware, gqlgin.serverHttpHandler(config.GqlServer))

	queryHandlers := append([]gin.HandlerFunc{
		gqlgenfn.GinContextToContextMiddleware(),
	}, config.Middleware...)

	rg.POST("/query", queryHandlers...)

	// only used for introspection, playground and protected by the ?pkey=xxx
	// also we do not use the /query for playground otherwise it unecessary runs the  PlaygroundAccessMiddleware
	// so we want to hyper optimize the speed and for /query we should not run this below middlewares to speed up
	introspectionHandlers := []gin.HandlerFunc{
		gqlgenfn.GinContextToContextMiddleware(),
		gqlgenfn.PlaygroundAccessMiddleware(config.PlaygroundKey),
		gqlgin.instrospectionMiddleware(config.GqlServer),
	}
	introspectionHandlers = append(introspectionHandlers, config.Middleware...)

	rg.POST("/query_playground", introspectionHandlers...)

	rg.GET("/gql", gqlgenfn.PlaygroundAccessMiddleware(config.PlaygroundKey), playgroundHandler(routeGroupPrefix))

	return gqlgin
}
