package appmiddleware

import (
	"context"
	"errors"
	"fmt"
	"lace/gqlgenfn"
	"net/http"
	"saas/gen/ent"
	"saas/gen/ent/app"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AppKey   = "app"
	AppError = "apperror"
)

// Get user from the gin gonic context actions
func GetAppFromGinCtx(c *gin.Context) (*ent.App, error) {
	// return the actual error not just generic error e.g. `access_denied` as setUserHandler is silent
	// for all the gql request so we will have to send the error in context
	userError, _ := c.Get(AppError)

	if userError != nil {
		return nil, fmt.Errorf(userError.(string))
	}

	record, _ := c.Get(AppKey)

	if record == nil {
		return nil, errors.New("user-not-found/invalid-token on auth context")
	}

	apRecord, ok := record.(*ent.App)
	if !ok {
		return nil, errors.New("something went wrong auth context")
	}

	return apRecord, nil
}

// this is for graphql only and will retrive the gin ctx first and the user from the gin ctx
// so we can get the current user info in the resolvers
// Ref: https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/gin.md
func GetAppFromGqlCtx(ctx context.Context) (*ent.App, error) {
	gc, err := gqlgenfn.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	appRecord, err := GetAppFromGinCtx(gc)

	if err != nil || appRecord == nil {
		return nil, err
	}
	return appRecord, nil
}

func MustGetAppFromGqlCtx(ctx context.Context) *ent.App {

	app, err := GetAppFromGqlCtx(ctx)
	if err != nil {
		panic(err)
	}

	return app
}

// get the token from query or header and set the user to context
func setAppHandler(c *gin.Context, client *ent.Client) error {

	appid := c.Query("app")

	xat := c.Request.Header.Get("x-app")
	if len(xat) > 0 {
		appid = xat
	}

	appid = strings.TrimSpace(appid)

	if len(appid) == 0 {
		return fmt.Errorf("no appid provided")
	}

	appRec, err := client.App.
		Query().
		Where(app.ID(appid)).
		Only(context.Background())

	if err != nil {
		// capture the error so we can pass the same to gql handler cuser, err := middleware.GetUserFromGqlCtx(ctx)
		// and show the actual error not just generic error i.e. `access_denied`
		c.Set(AppError, err.Error())

		return err
	}

	c.Set(AppKey, appRec)

	return nil
}

func CheckAppMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		isPlayground := gqlgenfn.IsPlaygroundGinCtx(c)
		if isPlayground {
			c.Next()
			return
		}

		err := setAppHandler(c, client)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{
					"message": err.Error(),
				},
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
