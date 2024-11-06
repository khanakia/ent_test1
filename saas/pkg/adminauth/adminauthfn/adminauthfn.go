package adminauthfn

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent"
	"saas/gen/ent/adminuser"
	"saas/pkg/auth/authfn"
	"strings"
)

func FindByEmail(email string, client *ent.Client) (*ent.AdminUser, error) {
	email = strings.ToLower(email)
	user, err := client.AdminUser.Query().Where(adminuser.Email(email)).First(context.Background())
	fmt.Println(err)

	if ent.IsNotFound(err) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return user, nil
}

func Login(params authfn.LoginParams, matchPass bool, client *ent.Client) (*ent.AdminUser, *ent.Session, error) {
	err := authfn.LoginValidate(params, matchPass)
	if err != nil {
		return nil, nil, err
	}

	user, err := FindByEmail(params.Email, client)
	if err != nil {
		return nil, nil, err
	}

	if !user.Status {
		return nil, nil, errors.New("user is disabled")
	}

	// if false we won't do the password matching and directly create the session
	if matchPass {
		matched := authfn.PasswordMatch(user.Password, params.Password)
		if !matched {
			return nil, nil, fmt.Errorf("password didn't match")
		}
	}

	session, err := authfn.CreateSession(authfn.CreateSessionParams{
		UserID:    user.ID,
		IP:        params.IP,
		UserAgent: params.UserAgent,
		Payload:   params.Payload,
		// ExpiresAt: time.Now().Add(8760 * time.Hour), // 1 year expiry
	}, client)

	if err != nil {
		return nil, nil, err
	}

	return user, session, nil
}
