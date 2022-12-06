package Tests

import (
	"CurrencyAPI/Helpers"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	userID := uint(1)

	hash, err := Helpers.GenerateToken(userID)

	assert.Equal(t, nil, err)
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(hash))

}
