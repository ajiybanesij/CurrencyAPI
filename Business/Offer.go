package Business

import (
	"CurrencyAPI/Config"
	"CurrencyAPI/Helpers"
	"CurrencyAPI/Repositories"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateOffer(fromCurrency, toCurrency string) (map[string]interface{}, error) {
	fromCurrency = strings.ToUpper(fromCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	if fromCurrency == toCurrency {
		return nil, errors.New("please select different currencies")
	}

	// Get Rate
	rate, ok := Config.RATES[fromCurrency+"-"+toCurrency]
	if !ok {
		return nil, errors.New("cuurency is invalid")
	}

	// From Wallet Check
	var fromWallet Repositories.Wallet
	fromCondition := map[string]interface{}{
		"currency": fromCurrency,
	}
	fromWalletError := fromWallet.Read(fromCondition)
	if fromWalletError != nil {
		fmt.Println(fromCurrency + " wallet not found")
		return nil, errors.New(fromCurrency + " wallet not found")
	}

	// To Wallet Check
	var toWallet Repositories.Wallet
	toCondition := map[string]interface{}{
		"currency": toCurrency,
	}
	toWalletError := toWallet.Read(toCondition)
	if toWalletError != nil {
		fmt.Println(toCurrency + " wallet not found")
		return nil, errors.New(toCurrency + " wallet not found")
	}

	// Create Offer Session
	id := uuid.New()
	offer := map[string]interface{}{
		"fromWalletID": fromWallet.ID,
		"fromCurrency": fromCurrency,
		"toWalletID":   toWallet.ID,
		"toCurrency":   toCurrency,
		"rate":         rate,
	}
	result, err := json.Marshal(offer)
	if err != nil {
		fmt.Println("Marshall Error")
		return nil, errors.New("error in creating offer")
	}

	// Save Offer Session
	expireTime := time.Now().Add(time.Minute * 3).Unix()
	Helpers.CreateRedis(id.String(), string(result))
	response := map[string]interface{}{
		"rate":    rate,
		"offerId": id.String(),
		"expire":  expireTime,
	}

	return response, nil
}

func AcceptOffer(offerID string, userId uint, amount float64) (*Repositories.Transaction, error) {
	// Offer Control
	offerInfoString, err := Helpers.ReadRedis(offerID)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("offer is invalid")
	}

	offerInfo := make(map[string]interface{})
	offerError := json.Unmarshal([]byte(*offerInfoString), &offerInfo)
	if offerError != nil {
		fmt.Println(offerError)
		return nil, errors.New("offer problem")
	}

	// Balance Control
	var balance Repositories.Balance
	balanceResult, err := balance.ReadBalance(offerInfo["fromWalletID"])
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("balance error")

	}

	walletBalance := *balanceResult
	rate := offerInfo["rate"].(float64)
	if amount/rate > walletBalance {
		fmt.Println("insufficient balance")
		return nil, errors.New("insufficient balance")
	}

	// Decrease Balance
	var fromBalance Repositories.Balance
	fromBalance.WalletID = uint(offerInfo["fromWalletID"].(float64))
	fromBalance.Amount = (-1) * (amount / rate)
	fromBalance.Currency = offerInfo["fromCurrency"].(string)

	fromBalanceErr := fromBalance.Create()
	if fromBalanceErr != nil {
		fmt.Println("balance decrease error", fromBalanceErr)
		return nil, fromBalanceErr
	}

	// Increase Balance
	var toBalance Repositories.Balance
	toBalance.WalletID = uint(offerInfo["toWalletID"].(float64))
	toBalance.Amount = (1) * (amount)
	toBalance.Currency = offerInfo["toCurrency"].(string)
	toBalanceErr := toBalance.Create()
	if toBalanceErr != nil {
		fmt.Println("balance decrease error", toBalanceErr)
		return nil, toBalanceErr
	}

	// Create Transaction
	var transaction Repositories.Transaction
	transaction.FromWalletID = fromBalance.WalletID
	transaction.FromCurrency = offerInfo["fromCurrency"].(string)
	transaction.FromAmount = fromBalance.Amount

	transaction.ToWalletID = toBalance.WalletID
	transaction.ToCurrency = offerInfo["toCurrency"].(string)
	transaction.ToAmount = toBalance.Amount

	transaction.Rate = rate
	transaction.UserID = userId

	transactionErr := transaction.Create()
	if transactionErr != nil {
		fmt.Println("transaction error", transactionErr)
		return nil, nil
	}

	redisDeleteErr := Helpers.DeleteRedis(offerID)
	if redisDeleteErr != nil {
		fmt.Println("offer can not deleted", redisDeleteErr)
		return &transaction, nil
	}

	return &transaction, nil
}
