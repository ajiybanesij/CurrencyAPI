package Tests

import (
	"CurrencyAPI/Helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	password := "testPassword123"

	hash, err := Helpers.HashPassword(password)
	match := Helpers.CheckPasswordHash(password, hash)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, match)

}
