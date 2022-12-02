package Models

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
