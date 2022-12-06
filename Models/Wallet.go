package Models

import "CurrencyAPI/Repositories"

type CreateWallet struct {
	Currency string `json:"currency" binding:"required"`
}

type ReadWallet struct {
	ID string `json:"id" binding:"required"`
}

type Wallet struct {
	ID          uint
	Currency    string
	Balance     float64
	Transaction *[]Repositories.Transaction
}
