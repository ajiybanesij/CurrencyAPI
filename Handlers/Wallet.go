package Handlers

import (
	"CurrencyAPI/Business"
	"CurrencyAPI/Helpers"
	"CurrencyAPI/Models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ReadCurrencies(ctx *gin.Context) {
	list := Business.ReadCurrencies()

	Helpers.Response(ctx, true, "success", list, 200)
	return
}

func ReadCurrencyRates(ctx *gin.Context) {
	list := Business.ReadCurrencyRates()

	Helpers.Response(ctx, true, "success", list, 200)
	return
}

func CreateWallet(ctx *gin.Context) {
	var walletModel Models.CreateWallet
	if err := ctx.ShouldBindJSON(&walletModel); err != nil {
		Helpers.Response(ctx, false, "body is not valid", nil, 400)
		return
	}
	userId, jwtErr := ctx.Get("userId")
	if !jwtErr {
		Helpers.Response(ctx, false, "jwt problem", nil, 400)
		return
	}

	result, err := Business.CreateWallet(userId.(uint), walletModel.Currency)
	if err != nil {
		Helpers.Response(ctx, false, err.Error(), nil, 400)
		return
	}
	Helpers.Response(ctx, true, "success", result, 201)
	return

}

func ReadWallet(ctx *gin.Context) {
	walletId := ctx.Param("id")
	if walletId == "" {
		Helpers.Response(ctx, false, "walletId required", nil, 400)
		return
	}
	userId, jwtErr := ctx.Get("userId")
	if !jwtErr {
		Helpers.Response(ctx, false, "jwt problem", nil, 400)
		return
	}

	u64, err := strconv.ParseUint(walletId, 10, 32)
	if err != nil {
		Helpers.Response(ctx, false, "walletId problem", nil, 400)
		return
	}
	walletID := uint(u64)
	result, err := Business.ReadWallet(userId.(uint), walletID)
	if err != nil {
		Helpers.Response(ctx, false, err.Error(), nil, 400)
		return
	}
	Helpers.Response(ctx, true, "success", result, 201)
	return
}

func ReadAllWallets(ctx *gin.Context) {
	userId, jwtErr := ctx.Get("userId")
	if !jwtErr {
		Helpers.Response(ctx, false, "jwt problem", nil, 400)
		return
	}

	result, err := Business.ReadWallets(userId.(uint))
	if err != nil {
		Helpers.Response(ctx, false, err.Error(), nil, 400)
		return
	}
	Helpers.Response(ctx, true, "success", result, 201)
	return
}
