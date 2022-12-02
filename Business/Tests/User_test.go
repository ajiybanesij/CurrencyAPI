package Tests

import (
	"CurrencyAPI/Business"
	"CurrencyAPI/Models"
	"CurrencyAPI/Repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNonExistUser(t *testing.T) {
	Repositories.InitDatabase()

	user := Models.User{
		Username: "nesij",
		Password: "hash",
	}
	err := Business.CreateUser(user)

	assert.Equal(t, "username already exists", err.Error())
}

func TestCreateUser(t *testing.T) {
	Repositories.InitDatabase()

	user := Models.User{
		Username: "nesij3",
		Password: "hash123",
	}

	err := Business.CreateUser(user)

	assert.Equal(t, nil, err)
}
