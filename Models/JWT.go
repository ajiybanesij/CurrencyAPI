package Models

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaim struct {
	UserId     uint  `json:"userId"`
	CreateTime int64 `json:"createTime"`
	Authorized bool  `json:"authorized"`
	jwt.StandardClaims
}
