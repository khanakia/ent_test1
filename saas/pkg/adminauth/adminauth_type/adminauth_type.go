package adminauth_type

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
	// UID   string    `json:"uid"`
	jwt.StandardClaims
}
