package gqldirectives

import (
	"context"
	"fmt"
	"saas/gen/ent/appuser"
	"saas/pkg/apppermfn"
	"saas/pkg/middleware/adminauthmiddleware"
	"saas/pkg/middleware/appmiddleware"

	"github.com/99designs/gqlgen/graphql"
)

func CanAdminDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	cuser, _ := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, fmt.Errorf("user not found")
	}

	if !cuser.Status {
		return nil, fmt.Errorf("access denied")
	}

	return next(ctx)
}

func CanAppDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	cuser, _ := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, fmt.Errorf("user not found")
	}

	if !cuser.Status {
		return nil, fmt.Errorf("access denied")
	}

	app, _ := appmiddleware.GetAppFromGqlCtx(ctx)
	if app == nil {
		return nil, fmt.Errorf("app not found")
	}

	appUser, err := app.QueryAppUsers().Where(appuser.AdminUserID(cuser.ID)).First(ctx)

	if err != nil {
		return nil, fmt.Errorf("app access denied")
	}

	perms, err := appUser.QueryAppRole().QueryAppRolePerms().All(ctx)

	if err != nil {
		return nil, fmt.Errorf("error fetcching permission")
	}

	// set permissions to ctx
	apppermfn.SetPermissionToGinCtx(ctx, perms)

	// or let it pass through
	return next(ctx)
}

func EnsureAppDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	app, _ := appmiddleware.GetAppFromGqlCtx(ctx)
	if app == nil {
		return nil, fmt.Errorf("app not found")
	}

	return next(ctx)
}
