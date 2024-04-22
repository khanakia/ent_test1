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

// Get user from the gin gonic context actions
func GetUserFromGinCtx(c *gin.Context) (*ent.User, error) {
	userc, _ := c.Get("user")

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
func GetUserFromGqlCtx(ctx context.Context) *ent.User {
	gc, err := gqlgenfn.GinContextFromContext(ctx)
	if err != nil {
		return nil
	}

	user, err := GetUserFromGinCtx(gc)

	if err != nil || user == nil {
		return nil
	}
	return user
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

	user, _, err := auth.ValidateJWT(token, client)

	if err != nil {
		return err
	}

	c.Set("user", user)

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
