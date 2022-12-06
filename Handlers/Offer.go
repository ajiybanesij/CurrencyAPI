package Handlers

import (
	"CurrencyAPI/Business"
	"CurrencyAPI/Helpers"
	"CurrencyAPI/Models"
	"github.com/gin-gonic/gin"
)

func CreateOffer(ctx *gin.Context) {
	var createOffer Models.CreateOffer
	if err := ctx.ShouldBindJSON(&createOffer); err != nil {
		Helpers.Response(ctx, false, "body is not valid", nil, 400)
		return
	}

	result, err := Business.CreateOffer(createOffer.FromCurrency, createOffer.ToCurrency)
	if err != nil {
		Helpers.Response(ctx, false, err.Error(), nil, 400)
		return
	}
	Helpers.Response(ctx, true, "success", result, 200)
	return
}

func AcceptOffer(ctx *gin.Context) {
	var acceptOffer Models.AcceptOffer
	if err := ctx.ShouldBindJSON(&acceptOffer); err != nil {
		Helpers.Response(ctx, false, "body is not valid", nil, 400)
		return
	}
	userId, jwtErr := ctx.Get("userId")
	if !jwtErr {
		Helpers.Response(ctx, false, "jwt problem", nil, 400)
		return
	}

	result, err := Business.AcceptOffer(acceptOffer.OfferID, userId.(uint), acceptOffer.Amount)
	if err != nil {
		Helpers.Response(ctx, false, err.Error(), nil, 400)
		return
	}
	Helpers.Response(ctx, true, "success", result, 200)
	return
}
