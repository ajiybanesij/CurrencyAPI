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
		Username: "nesij",
		Password: "testpass",
	}

	err := Business.CreateUser(user)

	assert.Equal(t, nil, err)
}

func TestLoginUser(t *testing.T) {
	Repositories.InitDatabase()

	user := Models.User{
		Username: "nesij",
		Password: "testpass",
	}

	jwt, err := Business.LoginUser(user)

	assert.NotEqual(t, nil, jwt)
	assert.Equal(t, nil, err)
}

func TestLoginUserWrongUsername(t *testing.T) {
	Repositories.InitDatabase()

	user := Models.User{
		Username: "nesijx",
		Password: "testpass",
	}

	_, err := Business.LoginUser(user)

	assert.Equal(t, "user is not exists", err.Error())
}

func TestLoginUserWrongPassword(t *testing.T) {
	Repositories.InitDatabase()

	user := Models.User{
		Username: "nesij",
		Password: "testpassX",
	}

	_, err := Business.LoginUser(user)

	assert.Equal(t, "invalid username or password", err.Error())
}
