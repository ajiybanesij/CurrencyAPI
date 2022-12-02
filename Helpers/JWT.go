package Helpers

import (
	"CurrencyAPI/Config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["expireTime"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Config.SECRET))
}
