package apppermfn

import (
	"context"
	"errors"
	"fmt"
	"lace/gqlgenfn"
	"saas/gen/ent"

	"github.com/gin-gonic/gin"
)

const (
	Perms      = "perms"
	PermError  = "permerror"
	Any        = "any"
	ReadAny    = "read_any"
	CreateAny  = "create_any"
	UpdateAny  = "update_any"
	DeleteAny  = "delete_any"
	ReadPost   = "read_post"
	CreatePost = "create_post"
	UpdatePost = "update_post"
	DeletePost = "delete_post"
)

type PermissionsMap map[string]bool

func GetPermsGinCtx(c *gin.Context) (PermissionsMap, error) {
	// return the actual error not just generic error e.g. `access_denied` as setUserHandler is silent
	// for all the gql request so we will have to send the error in context
	haserror, _ := c.Get(PermError)

	if haserror != nil {
		return nil, fmt.Errorf(haserror.(string))
	}

	permskey, _ := c.Get(Perms)

	if permskey == nil {
		return nil, errors.New("permission key not found")
	}

	result, ok := permskey.(PermissionsMap)
	if !ok {
		return nil, errors.New("something went wrong auth context")
	}

	return result, nil
}

func GetPermsFromGqlCtx(ctx context.Context) (PermissionsMap, error) {
	gc, err := gqlgenfn.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := GetPermsGinCtx(gc)

	if err != nil || user == nil {
		return nil, err
	}
	return user, nil
}

func SetPermissionToGinCtx(ctx context.Context, appPerms []*ent.AppRolePerm) error {

	gc, err := gqlgenfn.GinContextFromContext(ctx)
	if err != nil {
		return err
	}

	permsMap := PermissionsMap{}

	for _, perm := range appPerms {
		permsMap[perm.AppPermID] = true
	}

	gc.Set(Perms, permsMap)

	return nil
}

func can(ctx context.Context, permsSlice []string, appPermID string) error {
	permsMap, err := GetPermsFromGqlCtx(ctx)
	if err != nil {
		return err
	}

	// permsSlice := []string{Any, ReadAny, appPermID}

	for _, perm := range permsSlice {
		if _, ok := permsMap[perm]; ok {
			return nil
		}
	}

	return errors.New("access_denied")
}

func CanRead(ctx context.Context, appPermID string) error {
	permsSlice := []string{Any, ReadAny, appPermID}
	return can(ctx, permsSlice, appPermID)
}

func CanCreate(ctx context.Context, appPermID string) error {
	permsSlice := []string{Any, CreateAny, appPermID}
	return can(ctx, permsSlice, appPermID)
}

func CanUpdate(ctx context.Context, appPermID string) error {
	permsSlice := []string{Any, UpdateAny, appPermID}
	return can(ctx, permsSlice, appPermID)
}

func CanDelete(ctx context.Context, appPermID string) error {
	permsSlice := []string{Any, DeleteAny, appPermID}
	return can(ctx, permsSlice, appPermID)
}
