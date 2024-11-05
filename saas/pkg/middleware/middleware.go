package middleware

import (
	"context"
	"errors"
	"fmt"
	"lace/gqlgenfn"
	"net/http"
	"saas/gen/ent"
	"saas/pkg/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	User      = "user"
	UserError = "usererror"
)

// Get user from the gin gonic context actions
func GetUserFromGinCtx(c *gin.Context) (*ent.User, error) {
	// return the actual error not just generic error e.g. `access_denied` as setUserHandler is silent
	// for all the gql request so we will have to send the error in context
	userError, _ := c.Get(UserError)

	if userError != nil {
		return nil, fmt.Errorf(userError.(string))
	}

	userc, _ := c.Get(User)

	if userc == nil {
		return nil, errors.New("user-not-found/invalid-token on auth context")
	}

	user, ok := userc.(*ent.User)
	if !ok {
		return nil, errors.New("something went wrong auth context")
	}

	return user, nil
}

// this is for graphql only and will retrive the gin ctx first and the user from the gin ctx
// so we can get the current user info in the resolvers
// Ref: https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/gin.md
func GetUserFromGqlCtx(ctx context.Context) (*ent.User, error) {
	gc, err := gqlgenfn.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := GetUserFromGinCtx(gc)

	if err != nil || user == nil {
		return nil, err
	}
	return user, nil
}

// get the token from query or header and set the user to context
func setUserHandler(c *gin.Context, client *ent.Client) error {

	token := c.Query("token")

	xat := c.Request.Header.Get("token")
	if len(xat) > 0 {
		token = xat
	}

	token = strings.TrimSpace(token)

	if len(token) == 0 {
		return fmt.Errorf("no token provided")
	}

	user, err := auth.ValidateSession(token, client)

	if err != nil {
		// capture the error so we can pass the same to gql handler cuser, err := middleware.GetUserFromGqlCtx(ctx)
		// and show the actual error not just generic error i.e. `access_denied`
		c.Set(UserError, err.Error())

		return err
	}

	c.Set(User, user)

	return nil
}

// this middleware will always pass but we can use this to set some info let say
// curretn logged in usre information and then we can get it from ctx in gql resolvers
// take actions accordingly like access denied
func MiddlewareSilent(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		setUserHandler(c, client)

		c.Next()
	}
}

// for some routes we need to validate the authorization and fail if auth is not valid
func CheckAuthMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		isPlayground := gqlgenfn.IsPlaygroundGinCtx(c)

		if isPlayground {
			c.Next()
			return
		}

		err := setUserHandler(c, client)
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
