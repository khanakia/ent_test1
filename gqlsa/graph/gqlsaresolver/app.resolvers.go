package gqlsaresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"saas/gen/ent"
	"saas/gen/ent/app"
	"saas/pkg/middleware/adminauthmiddleware"
)

// CreateApp is the resolver for the createApp field.
func (r *mutationResolver) CreateApp(ctx context.Context, input ent.CreateAppInput) (*ent.App, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	return r.Plugin.EntDB.Client().App.Create().SetInput(input).SetAdminUserID(cuser.ID).Save(ctx)
}

// UpdateApp is the resolver for the updateApp field.
func (r *mutationResolver) UpdateApp(ctx context.Context, id string, input ent.UpdateAppInput) (*ent.App, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	return r.Plugin.EntDB.Client().App.UpdateOneID(id).SetInput(input).Where(app.AdminUserID(cuser.ID)).Save(ctx)
}
