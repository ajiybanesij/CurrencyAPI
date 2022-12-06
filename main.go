package main

import (
	"CurrencyAPI/Config"
	"CurrencyAPI/Repositories"
	"CurrencyAPI/Routes"
)

func main() {
	Repositories.InitDatabase()
	Config.InitRedis()

	r := Routes.SetupRouter()
	r.Run(":8081")
}
