package handlerfn

import (
	"context"
	"encoding/json"
	"fmt"
	"lace/util"
	"saas/gen/ent"
	"saas/gen/ent/temp"
	"saas/pkg/auth/authfn"
	"saas/pkg/emailfn"
	"saas/pkg/handler/handlertypes"
)

func RegisterHandler(input authfn.RegisterInput, client *ent.Client, ctx context.Context) error {
	err := authfn.RegisterValidate(input)
	if err != nil {
		return err
	}

	exists, err := authfn.CheckEmailExists(input.Email, client)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("email already registered")
	}

	tempRec, err := client.Temp.Create().
		SetAppID(input.AppID).
		SetType("register").
		SetBody(util.MustMarshalData(input)).
		Save(ctx)

	if err != nil {
		return err
	}

	// YTD - send email
	emailfn.RegisterVerify(input.AppID, input.Email, tempRec.ID, client)

	return err
}

func RegisterVerify(input handlertypes.RegisterVerifyInput, client *ent.Client, ctx context.Context) (*handlertypes.LoginResponse, error) {
	temp, err := client.Temp.Query().Where(temp.AppID(input.AppID), temp.ID(input.Token)).First(ctx)

	if err != nil {
		return nil, err
	}

	var registerInput authfn.RegisterInput

	err = json.Unmarshal(temp.Body, &registerInput)

	if err != nil {
		return nil, err
	}

	if input.Email != registerInput.Email {
		return nil, fmt.Errorf("email does not match")
	}

	user, err := authfn.RegisterWithTx(registerInput, client, ctx)

	if err != nil {
		return nil, err
	}

	client.Temp.DeleteOne(temp).Exec(ctx)

	// return user, err

	session, err := authfn.CreateSession(authfn.CreateSessionParams{
		AppID:  user.AppID,
		UserID: user.ID,
		// IP:        params.IP,
		// UserAgent: params.UserAgent,
		// Payload:   params.Payload,
	}, client)

	if err != nil {
		return nil, err
	}

	return &handlertypes.LoginResponse{
		Token: session.ID,
		Me:    user,
	}, nil
}
