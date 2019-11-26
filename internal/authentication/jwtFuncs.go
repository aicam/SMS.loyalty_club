package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create JWT token that is valid for 30 hours
func GenerateJWT(username string) (string, time.Time) {
	expirationTime := time.Now().Add(30 * time.Hour)
	expirationTimeUnix := expirationTime.Unix()
	claims := Claims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeUnix,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTKey())
	if err != nil {
		return "failed", expirationTime
	}
	return tokenString, expirationTime
}

func CheckToken(token string) int {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, e error) {
		return getJWTKey(), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0
		}
		return  -1
	}
	if !tkn.Valid {
		return 0
	}
	return 1
}