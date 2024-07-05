package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	env "github.com/mwdev22/ecom/app/config"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(user *User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(env.SecretKey)
}
