package gqlgenfn

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

type contextKey string

const (
	PlaygroundKey                = "pkey"
	isPlaygroundAllwedContextKey = "isplayground" //
)

func (c contextKey) String() string {
	return string(c)
}

// https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/gin.md
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKey("GinContextKey"), c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(contextKey("GinContextKey"))
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

// protect this route http://localhost:2303/sa/gql?pkey=1234
func PlaygroundAccessMiddleware(configkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		setPlaygroundContext(configkey, c)
		// var ok bool
		// key, ok := c.GetQuery(PlaygroundKey)
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "API Key required.",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// if len(key) == 0 || key != pkey {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "access denied.",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// c.Set(isPlaygroundAllwedContextKey, true)

		c.Next()
	}
}

func setPlaygroundContext(configkey string, gc *gin.Context) bool {
	key := gc.Query(PlaygroundKey)

	referer := gc.Request.Header.Get("Referer")
	if len(referer) != 0 && len(key) == 0 {
		u, err := url.Parse(referer)
		if err != nil {
			panic(err)
		}

		q := u.Query()
		key = q.Get(PlaygroundKey)
	}

	if len(key) == 0 {
		key = gc.Query(PlaygroundKey)
	}

	if key == configkey {
		gc.Set(isPlaygroundAllwedContextKey, true)
		return true
	}

	gc.Set(isPlaygroundAllwedContextKey, false)
	return false
}

// just check within gserver.AroundOperations and then disable the introspection accordingly
// func IsPlaygroundAllowedOperation(configkey string, ctx context.Context) bool {
// 	gc, err := GinContextFromContext(ctx)
// 	if err != nil {
// 		return false
// 	}

// 	return setPlaygroundContext(configkey, gc)

// 	// key := gc.Query(PlaygroundKey)

// 	// referer := gc.Request.Header.Get("Referer")
// 	// if len(referer) != 0 && len(key) == 0 {
// 	// 	u, err := url.Parse(referer)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}

// 	// 	q := u.Query()
// 	// 	key = q.Get(PlaygroundKey)
// 	// }

// 	// // fmt.Println("Key", key)
// 	// if len(key) == 0 {
// 	// 	key = gc.Query(PlaygroundKey)
// 	// }

// 	// if key == configkey {
// 	// 	gc.Set(isPlaygroundAllwedContextKey, true)
// 	// 	return true
// 	// }

// 	// gc.Set(isPlaygroundAllwedContextKey, false)
// 	// return false
// }

func IsPlaygroundGqlCtx(ctx context.Context) bool {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return false
	}

	return IsPlaygroundGinCtx(gc)
}

func IsPlaygroundGinCtx(c *gin.Context) bool {
	record, _ := c.Get(isPlaygroundAllwedContextKey)

	if record == nil {
		return false
	}

	value, _ := record.(bool)
	return value

}
