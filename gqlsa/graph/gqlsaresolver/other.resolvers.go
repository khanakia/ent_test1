package gqlsaresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"fmt"
	"saas/gen/ent"
	"saas/pkg/appfn"
	"saas/pkg/middleware"
)

// CreateOauthConnection is the resolver for the createOauthConnection field.
func (r *mutationResolver) CreateOauthConnection(ctx context.Context, input ent.CreateOauthConnectionInput) (*ent.OauthConnection, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	if !appfn.IsUserSA(cuser) {
		return nil, fmt.Errorf("unauthorized access")
	}
	return r.Plugin.EntDB.Client().OauthConnection.Create().SetInput(input).Save(ctx)
}

// UpdateOauthConnection is the resolver for the updateOauthConnection field.
func (r *mutationResolver) UpdateOauthConnection(ctx context.Context, id string, input ent.UpdateOauthConnectionInput) (*ent.OauthConnection, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	if !appfn.IsUserSA(cuser) {
		return nil, fmt.Errorf("unauthorized access")
	}

	return r.Plugin.EntDB.Client().OauthConnection.UpdateOneID(id).SetInput(input).Save(ctx)
}