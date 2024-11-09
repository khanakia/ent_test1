package gqldirectives

import (
	"context"
	"fmt"
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

	if app.AdminUserID != cuser.ID {
		return nil, fmt.Errorf("app access denied")
	}

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
