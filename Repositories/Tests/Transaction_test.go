package Tests

import (
	"CurrencyAPI/Repositories"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	Repositories.InitDatabase()

	var transaction Repositories.Transaction
	transaction.FromWalletID = 1
	transaction.FromCurrency = "EURO"
	transaction.FromAmount = -1

	transaction.ToWalletID = 1
	transaction.ToCurrency = "TRY"
	transaction.ToAmount = 19.63

	transaction.Rate = 19.63
	transaction.UserID = 1

	err := transaction.Create()
	assert.Equal(t, nil, err)
}

func TestReadTransaction(t *testing.T) {
	Repositories.InitDatabase()

	var transaction Repositories.Transaction

	condition := map[string]interface{}{
		"user_id": 1,
	}

	transactions, err := transaction.Read(condition)
	assert.Equal(t, nil, err)
	fmt.Println(transactions)
}
