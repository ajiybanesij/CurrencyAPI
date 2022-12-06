package Business

import (
	"CurrencyAPI/Config"
	"CurrencyAPI/Models"
	"CurrencyAPI/Repositories"
	"errors"
	"fmt"
	"strings"
)

func ReadCurrencies() []string {
	return Config.CURRENCIES
}
func ReadCurrencyRates() map[string]interface{} {
	return Config.RATES
}

func CreateWallet(userId uint, currency string) (*Models.Wallet, error) {
	var readWallet Repositories.Wallet
	condition := map[string]interface{}{
		"user_id":  userId,
		"currency": strings.ToUpper(currency),
	}
	walletReadError := readWallet.Read(condition)
	if walletReadError != nil && walletReadError.Error() != "record not found" {
		fmt.Println("wallet already exist", walletReadError)
		return nil, errors.New("wallet already exist")
	}
	if readWallet.Currency == strings.ToUpper(currency) {
		return nil, errors.New("wallet already exist")
	}

	var wallet Repositories.Wallet
	wallet.Currency = strings.ToUpper(currency)
	wallet.UserID = userId
	wallet.Balance = nil

	walletError := wallet.Create()
	if walletError != nil {
		fmt.Println("wallet create error", walletError)
		return nil, errors.New("wallet create error")
	}

	var balance Repositories.Balance
	balance.Amount = 1000.0
	balance.Currency = strings.ToUpper(currency)
	balance.WalletID = wallet.ID

	var response Models.Wallet
	response.ID = wallet.ID
	response.Currency = wallet.Currency

	balanceError := balance.Create()
	if balanceError != nil {
		fmt.Println("balance create error", balanceError)
		return &response, errors.New("wallet created but balance is 0")
	}
	response.Balance = 1000.0
	return &response, nil
}

func ReadWallet(userId uint, walletId uint) (*Models.Wallet, error) {
	var wallet Repositories.Wallet
	walletCondition := map[string]interface{}{
		"id":      walletId,
		"user_id": userId,
	}

	walletError := wallet.Read(walletCondition)
	if walletError != nil {
		fmt.Println("wallet read error", walletError)
		return nil, errors.New("wallet not found")
	}
	var balance Repositories.Balance

	balanceResult, err := balance.ReadBalance(wallet.ID)
	if err != nil {
		fmt.Println(err)
	}

	transactionCondition := map[string]interface{}{
		"from_wallet_id": walletId,
		"user_id":        userId,
	}
	var transaction Repositories.Transaction
	transactions, err := transaction.Read(transactionCondition)
	item := Models.Wallet{
		ID:          wallet.ID,
		Currency:    wallet.Currency,
		Balance:     *balanceResult,
		Transaction: transactions,
	}

	return &item, nil
}

func ReadWallets(userId uint) (*[]Models.Wallet, error) {
	var wallet Repositories.Wallet
	condition := map[string]interface{}{
		"user_id": userId,
	}

	wallets, walletError := wallet.ReadAll(condition)
	if walletError != nil {
		fmt.Println("wallet read error", walletError)
		return nil, errors.New("wallet not found")
	}

	var response []Models.Wallet
	for _, w := range wallets {
		var balance Repositories.Balance

		balanceResult, err := balance.ReadBalance(w.ID)
		if err != nil {
			fmt.Println(err)
		}

		item := Models.Wallet{
			ID:       w.ID,
			Currency: w.Currency,
			Balance:  *balanceResult,
		}
		response = append(response, item)
	}

	return &response, nil
}
