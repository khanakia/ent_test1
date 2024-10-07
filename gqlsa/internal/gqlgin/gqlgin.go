package gqlgin

import (
	"context"
	"fmt"
	"net/url"

	"lace/gqlgenfn"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type Gql struct {
	Config
}

type Config struct {
	RouterGroup      *gin.RouterGroup
	GqlServer        *handler.Server
	Middleware       []gin.HandlerFunc
	RouteGroupPrefix string
	PlaygroundKey    string
	// Resolver *resolverfn.Resolver
}

// Defining the Graphql handler
func (cls *Gql) graphqlHandler(gserver *handler.Server) gin.HandlerFunc {
	// https://github.com/99designs/gqlgen/blob/master/docs/content/reference/introspection.md
	gserver.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {

		isAllowedPlayground := IsPlaygroundAllwedForContext(cls.PlaygroundKey, ctx)
		fmt.Println("isAllowedPlayground", isAllowedPlayground)
		if !isAllowedPlayground {
			graphql.GetOperationContext(ctx).DisableIntrospection = true
		}

		return next(ctx)
	})

	return func(c *gin.Context) {
		gserver.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler(routeGroupPrefix string) gin.HandlerFunc {
	h := playground.Handler("GraphQL", fmt.Sprintf("%s/query", routeGroupPrefix))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func IsPlaygroundAllwedForContext(configkey string, ctx context.Context) bool {
	gc, err := gqlgenfn.GinContextFromContext(ctx)
	if err != nil {
		return false
	}

	key := gc.Query("key")

	referer := gc.Request.Header.Get("Referer")
	if len(referer) != 0 && len(key) == 0 {
		u, err := url.Parse(referer)
		if err != nil {
			panic(err)
		}

		q := u.Query()
		key = q.Get("key")
	}

	// fmt.Println("Key", key)
	if len(key) == 0 {
		key = gc.Query("key")
	}

	if key == configkey {
		return true
	}

	return false
}

func playgroundAccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var ok bool
		// key, ok := c.GetQuery("key")
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "API Key required.",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// if key != viper.GetString("gql_playground_key") {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "Wrong key.",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		c.Set("allow_playground", true)

		c.Next()
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

	queryHandlers := append(config.Middleware, gqlgin.graphqlHandler(config.GqlServer))
	rg.POST("/query", queryHandlers...)

	rg.GET("/gql", playgroundAccessMiddleware(), playgroundHandler(routeGroupPrefix))

	return gqlgin
}
