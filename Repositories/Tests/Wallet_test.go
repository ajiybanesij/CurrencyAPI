package Tests

import (
	"CurrencyAPI/Repositories"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	Repositories.InitDatabase()

	var wallet Repositories.Wallet
	wallet.UserID = 2
	wallet.Currency = "EURO"

	err := wallet.Create()
	assert.Equal(t, nil, err)
}

func TestReadWallet(t *testing.T) {
	Repositories.InitDatabase()

	var wallet Repositories.Wallet
	condition := map[string]interface{}{
		"id": 3,
	}

	err := wallet.Read(condition)
	assert.Equal(t, nil, err)
	fmt.Println(wallet)
}

func TestReadAllWallets(t *testing.T) {
	Repositories.InitDatabase()

	var wallet Repositories.Wallet
	condition := map[string]interface{}{
		"user_id": 2,
	}

	wallets, err := wallet.ReadAll(condition)
	assert.Equal(t, nil, err)
	fmt.Println(wallets)
}
