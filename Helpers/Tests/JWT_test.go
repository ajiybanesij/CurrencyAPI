package Tests

import (
	"CurrencyAPI/Helpers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	username := "nesij"

	hash, err := Helpers.GenerateToken(username)

	assert.Equal(t, nil, err)
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(hash))

}
