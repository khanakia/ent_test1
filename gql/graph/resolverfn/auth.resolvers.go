package resolverfn

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"gql/graph/model"
	"lace/util"
	"saas/gen/ent"
	"saas/pkg/auth/authfn"
	"saas/pkg/handler/handlerfn"
	"saas/pkg/handler/handlertypes"
	"saas/pkg/middleware"

	"github.com/spf13/cast"
)

// AuthLoginViaOauth is the resolver for the authLoginViaOauth field.
func (r *mutationResolver) AuthLoginViaOauth(ctx context.Context, cacheID string) (*handlertypes.LoginResponse, error) {
	panic(fmt.Errorf("not implemented: AuthLoginViaOauth - authLoginViaOauth"))
}

// AuthRegister is the resolver for the authRegister field.
func (r *mutationResolver) AuthRegister(ctx context.Context, input authfn.RegisterInput) (bool, error) {
	err := handlerfn.RegisterHandler(authfn.RegisterInput{
		Email:     input.Email,
		Password:  input.Password,
		Phone:     cast.ToString(input.Phone),
		FirstName: cast.ToString(input.FirstName),
		LastName:  cast.ToString(input.LastName),
		Company:   cast.ToString(input.Company),
	}, r.Plugin.EntDB.Client(), ctx)

	if err != nil {
		return false, util.ErrorToGqlError(err, ctx)
	}

	return true, err
}

// AuthRegisterVerify is the resolver for the authRegisterVerify field.
func (r *mutationResolver) AuthRegisterVerify(ctx context.Context, input handlertypes.RegisterVerifyInput) (*handlertypes.LoginResponse, error) {
	result, err := handlerfn.RegisterVerify(handlertypes.RegisterVerifyInput{
		Email: input.Email,
		Token: input.Token,
	}, r.Plugin.EntDB.Client(), ctx)

	if err != nil {
		return nil, util.ErrorToGqlError(err, ctx)
	}

	return result, err
}

// AuthLogin is the resolver for the authLogin field.
func (r *mutationResolver) AuthLogin(ctx context.Context, input authfn.LoginParams) (*handlertypes.LoginResponse, error) {
	user, session, err := authfn.Login(input, true, r.Plugin.EntDB.Client())

	if err != nil {
		return nil, util.ErrorToGqlError(err, ctx)
	}

	return &handlertypes.LoginResponse{
		Token: session.ID,
		Me:    user,
	}, err
}

// AuthForgotPassword is the resolver for the authForgotPassword field.
func (r *mutationResolver) AuthForgotPassword(ctx context.Context, userName string) (bool, error) {
	err := authfn.ForgotPassword(userName, r.Plugin.EntDB.Client(), r.Plugin.Cache)

	if err != nil {
		return false, err
	}

	return true, nil
}

// AuthResetPassword is the resolver for the authResetPassword field.
func (r *mutationResolver) AuthResetPassword(ctx context.Context, input model.ResetPasswordInput) (bool, error) {
	_, err := authfn.ResetPassword(authfn.ResetPasswordParams{
		Email:           input.Email,
		Token:           input.Token,
		Password:        input.Password,
		PasswordConfirm: input.PasswordConfirm,
	}, r.Plugin.EntDB.Client(), r.Plugin.Cache)

	if err != nil {
		return false, err
	}

	return true, nil
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (bool, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return false, err
	}

	if input.Password != input.PasswordConfirm {
		return false, fmt.Errorf("passwords don't match")
	}

	matched := authfn.PasswordMatch(cuser.Password, input.OldPassword)
	if !matched {
		return false, fmt.Errorf("password didn't match")
	}

	_, err = cuser.Update().SetPassword(authfn.Hash(input.Password)).Save(ctx)

	if err != nil {
		return false, fmt.Errorf("something went wrong")
	}

	return true, nil
}

// UpdateProfile is the resolver for the updateProfile field.
func (r *mutationResolver) UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (bool, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return false, err
	}

	_, err = cuser.Update().
		SetFirstName(cast.ToString(input.FirstName)).
		SetLastName(cast.ToString(input.LastName)).
		SetCompany(cast.ToString(input.Company)).
		SetPhone(cast.ToString(input.Phone)).
		Save(ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	cuser, err := middleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	return cuser, nil
}
