package authfn

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"lace/cache"
	"saas/gen/ent"
	"saas/gen/ent/user"
	"saas/pkg/appfn"
	"saas/pkg/emailfn"
	"strings"
	"time"

	"github.com/cohesivestack/valgo"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash)
}

// PasswordMatch - Compare two passwords are equal
func PasswordMatch(password string, password1 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password1))
	return err == nil
}

func CheckEmailExists(email string, client *ent.Client) (bool, error) {
	count, err := client.User.Query().Where(user.Email(email)).Count(context.Background())
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}

func FindByEmail(email string, client *ent.Client) (*ent.User, error) {
	email = strings.ToLower(email)
	user, err := client.User.Query().Where(user.Email(email)).First(context.Background())
	fmt.Println(err)

	if ent.IsNotFound(err) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return user, nil
}

type RegisterInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Company   string `json:"company,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Password  string `json:"password"`
}

func RegisterValidate(userInput RegisterInput) error {
	val := valgo.Is(
		valgo.String(userInput.Email, "email").Not().Blank().OfLengthBetween(4, 20),
		valgo.String(userInput.Password, "password").Not().Blank().OfLengthBetween(4, 20),
		// valgo.Number(17, "age").GreaterThan(18),
	)
	if !val.Valid() {
		out, _ := json.MarshalIndent(val.Error(), "", "  ")
		fmt.Println(string(out))
		return val.Error()
	}
	return nil
}

func RegisterWithTx(userInput RegisterInput, client *ent.Client, ctx context.Context) (*ent.User, error) {
	return appfn.WithTx(ctx, client, func(tx *ent.Tx) (*ent.User, error) {
		return Register(userInput, tx.Client(), ctx)
	})
}

func Register(userInput RegisterInput, client *ent.Client, ctx context.Context) (*ent.User, error) {
	err := RegisterValidate(userInput)
	if err != nil {
		return nil, err
	}

	exists, err := CheckEmailExists(userInput.Email, client)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("email already registered")
	}

	hashed := (Hash(userInput.Password))

	creator := client.User.Create().
		SetEmail(userInput.Email).
		SetFirstName(cast.ToString(userInput.FirstName)).
		SetLastName(cast.ToString(userInput.LastName)).
		SetPassword(hashed).
		SetStatus(true).
		SetWelcomeEmailSent(false)

	user, err := creator.Save(ctx)
	if err != nil {
		return nil, err
	}

	err = appfn.CreateDefaultWorkspaceForUser(user.ID, client, ctx)

	if err != nil {
		return nil, err
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

func Login(params LoginParams, matchPass bool, client *ent.Client) (*ent.User, *ent.Session, error) {
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
		matched := PasswordMatch(user.Password, params.Password)
		if !matched {
			return nil, nil, fmt.Errorf("password didn't match")
		}
	}

	session, err := CreateSession(CreateSessionParams{
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

type CreateSessionParams struct {
	UserID    string
	IP        string
	UserAgent string
	Payload   string
	ExpiresAt time.Time
}

func CreateSession(params CreateSessionParams, client *ent.Client) (*ent.Session, error) {

	if params.ExpiresAt.IsZero() {
		params.ExpiresAt = time.Now().Add(8760 * time.Hour) // 1 year expiry
	}

	session, err := client.Session.Create().
		SetUserID(params.UserID).
		SetUserAgent(params.UserAgent).
		SetIP(params.IP).
		SetPayload(params.Payload).
		SetExpiresAt(params.ExpiresAt).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return session, nil
}

func ForgotPassword(email string, client *ent.Client, cache cache.Store) error {

	if len(email) == 0 {
		return fmt.Errorf("email must not be empty")
	}

	user, err := FindByEmail(email, client)
	if err != nil {
		return err
	}

	if !user.Status {
		return errors.New("user is disabled")
	}

	secret := gonanoid.Must()
	token := "fp:" + secret

	_, err = cache.Put(token, user.ID, 1000)

	if err != nil {
		return err
	}

	emailfn.SendForgotPassword(email, secret, client)

	return nil
}

type ResetPasswordParams struct {
	Email           string `json:"email"`
	Token           string `json:"token"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func ResetPassword(params ResetPasswordParams, client *ent.Client, cache cache.Store) (*ent.User, error) {

	if params.Password != params.PasswordConfirm {
		return nil, fmt.Errorf("passwords don't match")
	}

	cacheKey := "fp:" + params.Token
	userID := cache.Get(cacheKey)

	if userID == nil {
		return nil, errors.New("wrong token supplied")
	}

	user, err := client.User.Get(context.Background(), userID.(string))
	if err != nil {
		return nil, err
	}

	if user.Email != params.Email {
		return nil, errors.New("wrong email supplied")
	}

	user, err = user.Update().SetPassword(Hash(params.Password)).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	cache.Del(cacheKey)

	return user, nil
}
