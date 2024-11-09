package gqlsaresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"gqlsa/graph/model"
	"lace/util"
	"saas/pkg/adminauth/adminauthfn"
)

// AdminAuthLogin is the resolver for the adminAuthLogin field.
func (r *mutationResolver) AdminAuthLogin(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	user, session, err := adminauthfn.Login(adminauthfn.LoginParams{
		Email:    input.Email,
		Password: input.Password,
	}, true, r.Plugin.EntDB.Client())

	if err != nil {
		return nil, util.ErrorToGqlError(err, ctx)
	}

	return &model.LoginResponse{
		Token: session.ID,
		Me:    user,
	}, err
}