package Helpers

import (
	"CurrencyAPI/Config"
	"CurrencyAPI/Models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userID
	claims["createTime"] = time.Now().Unix()
	//claims["expireTime"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Config.SECRET))
}

func ValidateToken(signedToken string) (*Models.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(Config.SECRET), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Models.JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	return claims, err
}
