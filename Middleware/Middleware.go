package Middleware

import (
	"CurrencyAPI/Helpers"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(ctx *gin.Context) {
	jwt := ctx.Request.Header["Authorization"]
	if jwt == nil {
		Helpers.Response(ctx, false, "jwt required", nil, 401)
		ctx.Abort()
		return
	}

	baseToken := strings.Split(jwt[0], " ")

	jwtResult, err := Helpers.ValidateToken(baseToken[1])
	if err != nil {
		Helpers.Response(ctx, false, "jwt is invalid", nil, 401)
		ctx.Abort()
		return
	}
	ctx.Set("userId", jwtResult.UserId)
	ctx.Next()
}
