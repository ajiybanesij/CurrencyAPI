package main

import (
	"CurrencyAPI/Business"
	"CurrencyAPI/Repositories"
)

func main() {
	Repositories.InitDatabase()

	Business.CreateUser()
}
