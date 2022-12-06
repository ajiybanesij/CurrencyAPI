package Helpers

import (
	"CurrencyAPI/Config"
	"context"
	"time"
)

var ctx = context.Background()

func CreateRedis(key string, value string) error {
	if err := Config.GetRedisInstance().Set(ctx, key, value, 180*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func ReadRedis(key string) (*string, error) {
	val, err := Config.GetRedisInstance().Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func DeleteRedis(key string) error {
	if err := Config.GetRedisInstance().Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
