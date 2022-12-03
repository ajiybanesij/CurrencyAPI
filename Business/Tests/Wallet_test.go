package Tests

import (
	"CurrencyAPI/Business"
	"CurrencyAPI/Repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	Repositories.InitDatabase()

	userID := uint(13)
	currency := "TRY"

	wallet, err := Business.CreateWallet(userID, currency)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, wallet.ID)
}

func TestReadWallet(t *testing.T) {
	Repositories.InitDatabase()

	userID := uint(13)
	walletID := uint(2)

	wallet, err := Business.ReadWallet(userID, walletID)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, wallet.ID)
}

func TestReadWallets(t *testing.T) {
	Repositories.InitDatabase()

	userID := uint(13)

	_, err := Business.ReadWallets(userID)
	assert.Equal(t, nil, err)
}
