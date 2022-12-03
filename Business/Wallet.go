package Business

import (
	"CurrencyAPI/Repositories"
	"errors"
	"fmt"
)

func CreateWallet(userId uint, currency string) (*Repositories.Wallet, error) {
	var wallet Repositories.Wallet
	wallet.Currency = currency
	wallet.UserID = userId
	wallet.Balance = nil

	walletError := wallet.Create()
	if walletError != nil {
		fmt.Println("wallet create error", walletError)
		return nil, errors.New("wallet create error")
	}

	var balance Repositories.Balance
	balance.Amount = 1000.0
	balance.Currency = currency
	balance.WalletID = wallet.ID

	balanceError := balance.Create()
	if balanceError != nil {
		fmt.Println("balance create error", balanceError)
		return &wallet, errors.New("wallet created but balance is 0")
	}
	return &wallet, nil
}

func ReadWallet(userId uint, walletId uint) (*Repositories.Wallet, error) {
	var wallet Repositories.Wallet
	condition := map[string]interface{}{
		"id":      walletId,
		"user_id": userId,
	}

	walletError := wallet.Read(condition)
	if walletError != nil {
		fmt.Println("wallet read error", walletError)
		return nil, errors.New("wallet not found")
	}

	return &wallet, nil
}

func ReadWallets(userId uint) (*[]Repositories.Wallet, error) {
	var wallet Repositories.Wallet
	condition := map[string]interface{}{
		"user_id": userId,
	}

	wallets, walletError := wallet.ReadAll(condition)
	if walletError != nil {
		fmt.Println("wallet read error", walletError)
		return nil, errors.New("wallet not found")
	}

	return &wallets, nil
}
