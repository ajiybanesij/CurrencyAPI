package Tests

import (
	"CurrencyAPI/Business"
	"CurrencyAPI/Config"
	"CurrencyAPI/Repositories"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateOffer(t *testing.T) {
	Repositories.InitDatabase()
	Config.InitRedis()

	from := "TRY"
	to := "EURO"

	result, err := Business.CreateOffer(from, to)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, result)
	fmt.Println(result)

}

func TestAcceptOfferFail(t *testing.T) {
	Repositories.InitDatabase()
	Config.InitRedis()

	offerId := "4cc2406e-b34a-4ff1-81c4-7d5c0c1e6a20"
	resp, err := Business.AcceptOffer(offerId, 10)
	assert.Equal(t, nil, resp)
	assert.NotEqual(t, nil, err)
}

func TestOffer(t *testing.T) {
	Repositories.InitDatabase()
	Config.InitRedis()

	from := "TRY"
	to := "EURO"

	result, err := Business.CreateOffer(from, to)

	offerId := result["offerId"].(string)
	resp, err := Business.AcceptOffer(offerId, 10)

	assert.NotEqual(t, nil, resp)
	assert.Equal(t, nil, err)

	fmt.Println(resp)

}
