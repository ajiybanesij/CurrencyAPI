package Handlers

import (
	"CurrencyAPI/Business"
	"CurrencyAPI/Helpers"
	"CurrencyAPI/Models"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user Models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		Helpers.Response(ctx, false, "body is not valid", nil, 400)
		return
	}

	err := Business.CreateUser(user)
	if err != nil {
		Helpers.Response(ctx, false, err.Error(), nil, 400)
		return
	}
	Helpers.Response(ctx, true, "success", nil, 201)
	return
}

func LoginUser(ctx *gin.Context) {
	var user Models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		Helpers.Response(ctx, false, "body is not valid", nil, 400)
		return
	}

	jwt, err := Business.LoginUser(user)
	if err != nil {
		Helpers.Response(ctx, false, err.Error(), nil, 400)
		return
	}
	Helpers.Response(ctx, true, "success", jwt, 200)
	return

}
