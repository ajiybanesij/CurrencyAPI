package Tests

import (
	"CurrencyAPI/Config"
	"CurrencyAPI/Helpers"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisCreate(t *testing.T) {
	Config.InitRedis()

	condition := map[string]interface{}{
		"test": "test",
	}
	value, _ := json.Marshal(condition)

	err := Helpers.CreateRedis("test", string(value))
	assert.Equal(t, nil, err)
	time.Sleep(5 * time.Second)
	result, err := Helpers.ReadRedis("test")

	assert.NotEqual(t, value, result)
}
