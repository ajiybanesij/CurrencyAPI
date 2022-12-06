package Tests

import (
	"CurrencyAPI/Repositories"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBalance(t *testing.T) {
	Repositories.InitDatabase()

	var balance Repositories.Balance
	balance.WalletID = 1
	balance.Amount = -10

	err := balance.Create()
	assert.Equal(t, nil, err)
}

func TestReadBalance(t *testing.T) {
	Repositories.InitDatabase()
	_ = map[string]interface{}{
		"wallet_id": uint(3),
	}
	var balance Repositories.Balance
	sum, err := balance.ReadBalance(uint(1))
	assert.Equal(t, nil, err)
	fmt.Println(*sum)
}
