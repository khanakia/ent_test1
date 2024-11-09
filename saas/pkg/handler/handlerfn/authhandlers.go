package handlerfn

import (
	"context"
	"encoding/json"
	"fmt"
	"lace/cache"
	"lace/publicid"
	"lace/util"
	"saas/gen/ent"
	"saas/gen/ent/temp"
	"saas/pkg/auth/authfn"
	"saas/pkg/emailfn"
	"saas/pkg/handler/handlertypes"
	"saas/pkg/oauth/oauthconnectionfn"
	"saas/pkg/oauth/oauthuserinfo"

	"github.com/spf13/cast"
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

	// session, err := authfn.CreateSession(authfn.CreateSessionParams{
	// 	AppID:  user.AppID,
	// 	UserID: user.ID,
	// 	// IP:        params.IP,
	// 	// UserAgent: params.UserAgent,
	// 	// Payload:   params.Payload,
	// }, client)

	user, session, err := authfn.Login(authfn.LoginParams{
		AppID: input.AppID,
		Email: user.Email,
		// IP        string
		// UserAgent string
		// Payload   string
	}, false, client)

	if err != nil {
		return nil, err
	}

	return &handlertypes.LoginResponse{
		Token: session.ID,
		Me:    user,
	}, nil
}

func OauthLoginOrRegisterHandler(ctx context.Context, appID, cacheID string, allowRegister bool, client *ent.Client, cache cache.Store) (*ent.User, *ent.Session, error) {
	oauthRespCacheAny := cache.Get(cacheID)

	var oauthRespCache oauthconnectionfn.OauthResponseCache
	err := util.InterfaceToStruct(oauthRespCacheAny, &oauthRespCache)

	if err != nil {
		fmt.Println(err)
		return nil, nil, fmt.Errorf("something went wrong")
	}

	if oauthRespCache.AppID != appID {
		return nil, nil, fmt.Errorf("access denied app")
	}

	oauth2Token, err := oauthconnectionfn.OauthTokenFromResponseCache(oauthRespCache)

	if err != nil {
		fmt.Println(err)
		return nil, nil, fmt.Errorf("something went wrong")
	}

	userinfo, err := oauthuserinfo.UserInfoGet(oauthRespCache.Provider, oauth2Token)

	if err != nil {
		fmt.Println(err)
		return nil, nil, fmt.Errorf("something went wrong")
	}

	// check if user is not in database
	isNew := false
	if allowRegister {
		_, err := authfn.FindByEmail(appID, userinfo.Email, client)

		if ent.IsNotFound(err) {
			isNew = true
		}

		if err != nil && !ent.IsNotFound(err) {
			return nil, nil, err
		}
	}

	// if new then register the user and signin immediately
	if isNew {
		_, err := authfn.RegisterWithTx(authfn.RegisterInput{
			AppID:     appID,
			Email:     userinfo.Email,
			Password:  publicid.Must(),
			FirstName: cast.ToString(userinfo.FirstName),
			LastName:  cast.ToString(userinfo.LastName),
		}, client, ctx)

		if err != nil {
			return nil, nil, err
		}

		user, session, err := authfn.Login(authfn.LoginParams{
			AppID: appID,
			Email: userinfo.Email,
			// IP        string
			// UserAgent string
			// Payload   string
		}, false, client)

		return user, session, err
	}

	// if not new the try to login
	user, session, err := authfn.Login(authfn.LoginParams{
		AppID: appID,
		Email: userinfo.Email,
	}, false, client)

	if err != nil {
		return nil, nil, err
	}

	cache.Del(cacheID)

	return user, session, err
}
