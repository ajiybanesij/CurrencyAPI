package Tests

import (
	"CurrencyAPI/Repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	Repositories.InitDatabase()

	user := Repositories.User{
		Name: "Nesij" + uuid.New().String(),
		Hash: "testHash" + uuid.New().String(),
	}
	err := user.Create()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadUserSuccess(t *testing.T) {
	Repositories.InitDatabase()
	var user Repositories.User
	query := map[string]interface{}{"name": "Nesij"}

	err := user.Read(query)

	assert.Equal(t, "testHash", user.Hash)
	assert.Equal(t, nil, err)
}

func TestReadUserFail(t *testing.T) {
	Repositories.InitDatabase()
	var user Repositories.User
	query := map[string]interface{}{"name": "Nesij2"}

	err := user.Read(query)

	assert.NotEqual(t, "testHash", user.Hash)
	assert.NotEqual(t, nil, err)
}
