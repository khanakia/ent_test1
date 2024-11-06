package adminauthfn

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent"
	"saas/gen/ent/adminuser"
	"saas/pkg/auth/authfn"
	"strings"

	"github.com/cohesivestack/valgo"
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

type LoginParams struct {
	Email     string
	Password  string
	IP        string
	UserAgent string
	Payload   string
	// ExpiresAt time.Time
}

func LoginValidate(params LoginParams, matchPass bool) error {
	val := valgo.Is(
		valgo.String(params.Email, "email").Not().Blank().OfLengthBetween(4, 20),
	)

	if matchPass {
		val.Is(
			valgo.String(params.Password, "password").Not().Blank().OfLengthBetween(4, 20),
		)
	}

	if !val.Valid() {
		return val.Error()
	}

	return nil
}

func Login(params LoginParams, matchPass bool, client *ent.Client) (*ent.AdminUser, *ent.Session, error) {
	err := LoginValidate(params, matchPass)
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
		AppID:     "admin",
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
