package auth_type

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Shop  string `json:"shop"`
	Name  string `json:"name"`
	// UID   string    `json:"uid"`
	jwt.StandardClaims
}

type UserShopifySetting struct {
	AppKey    string
	AppSecret string
}
