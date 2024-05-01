package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"lace/publicid"
	"saas/gen/ent"
	"saas/gen/ent/session"
	"saas/gen/ent/user"
	"saas/pkg/auth/auth_type"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/ubgo/goutil"
)

func GenerateUID() string {
	secret := uuid.New().String()
	return secret
}

func GetUserSecret(user ent.User, client *ent.Client) string {
	if len(user.Secret) == 0 {
		user.Update().SetSecret(publicid.Must()).Save(context.Background())
		return user.Secret
	}
	return user.Secret
}

/*
 * This signature is used to create a JWT token
 * Benefits - By adding the user secret with the appsecret it makes the app more secure
 * Let say if appSecret is compromised then stil nobody can generate tokens without user secret
 * If userSecret compromised it will affect only single user not all the users
 * If user wants to focefully logout for all the applications we simply update his userSecret
 * FUTURE CONSIDERATION - Add jwt to token to the blacklist if users logout
 */
func GetSignature(user ent.User, client *ent.Client) string {
	userSecret := GetUserSecret(user, client)

	appSecret := goutil.Env("appSecret", "IBIrewORShiVReBASTer")
	signature := appSecret + ":" + userSecret
	return signature
}

func CreateToken(user ent.User, client *ent.Client) (string, error) {
	// expirationTime := time.Now().Add(500 * time.Minute) // 500 minute

	// expire after 30 days
	expires := time.Now().AddDate(0, 0, 30)

	claims := &auth_type.Claims{
		Sub:   user.ID,
		Email: user.Email,
		Name:  user.FirstName + " " + user.LastName,
		// ID:    user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expires.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signature := []byte(GetSignature(user, client))
	tokenString, err := token.SignedString(signature)

	return tokenString, err
}

func ValidateJWT(token string, client *ent.Client) (*ent.User, *auth_type.Claims, error) {
	token = strings.TrimSpace(token)

	claims := &auth_type.Claims{}

	_, _, err := new(jwt.Parser).ParseUnverified(token, claims)
	if err != nil {
		return nil, nil, err
	}

	if len(claims.Sub) == 0 {
		return nil, nil, errors.New("userID is empty")
	}

	userRec, err := client.User.
		Query().
		Where(user.ID(claims.Sub)).
		Only(context.Background())

	if err != nil && ent.IsNotFound(err) {
		return nil, nil, errors.New("User not found")
	}

	if err != nil {
		return nil, nil, err
	}

	if !userRec.Status {
		return nil, nil, errors.New("user is disabled")
	}

	signature := []byte(GetSignature(*userRec, client))

	tokenWC, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return signature, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !tokenWC.Valid {
		return nil, nil, errors.New("Unauthorized")
	}

	return userRec, claims, nil
}

func ValidateSession(sessionId string, client *ent.Client) (*ent.User, error) {
	session, err := client.Session.Query().Where(session.ID(sessionId)).First(context.Background())

	if err != nil {
		return nil, fmt.Errorf("session not found")
	}

	userRec, err := client.User.
		Query().
		Where(user.ID(session.UserID)).
		Only(context.Background())

	if err != nil && ent.IsNotFound(err) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	if !userRec.Status {
		return nil, errors.New("user is disabled")
	}

	return userRec, nil
}

func GetUserFromContext(c *gin.Context) (*ent.User, error) {
	userc, _ := c.Get("user")

	if userc == nil {
		return nil, errors.New("User not found")
	}
	user, _ := userc.(*ent.User)
	return user, nil
}

type UserShopifySetting struct {
	AppKey    string
	AppSecret string
	AppHandle string
}

func UserFillDefaults(mutation *ent.UserMutation) {
	if _, ok := mutation.Status(); !ok {
		mutation.SetStatus(true)
	}

	if _, ok := mutation.Password(); !ok {
		randomPass := publicid.Must()
		mutation.SetPassword(randomPass)
	}

	if _, ok := mutation.Secret(); !ok {
		mutation.SetSecret(uuid.New().String())
	}

	if _, ok := mutation.APIKey(); !ok {
		mutation.SetAPIKey(uuid.New().String())
	}
}

func IsValidHmac(message string, hmacOld string, secret string) bool {
	// mac := hmac.New(sha256.New, []byte(viper.GetString("shopify.client_secret")))
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	fmt.Printf("\nHMAC-Sha256: %x \n", mac.Sum(nil))
	hashInBytes := mac.Sum(nil)
	hashInString := hex.EncodeToString(hashInBytes[:])
	isValidMac := (hashInString == hmacOld)

	return isValidMac
}
