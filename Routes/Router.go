package Routes

import (
	"CurrencyAPI/Handlers"
	"CurrencyAPI/Middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/", Middleware.Auth)
	unAuth := router.Group("/")

	unAuth.POST("/register", Handlers.CreateUser)
	unAuth.POST("/login", Handlers.LoginUser)

	auth.GET("/currency/all", Handlers.ReadCurrencies)
	auth.GET("/currency/rates", Handlers.ReadCurrencyRates)

	auth.POST("/wallet", Handlers.CreateWallet)
	auth.GET("/wallet/:id", Handlers.ReadWallet)
	auth.GET("/wallet", Handlers.ReadAllWallets)

	auth.POST("/offer/create", Handlers.CreateOffer)
	auth.POST("/offer/accept", Handlers.AcceptOffer)

	return router
}
