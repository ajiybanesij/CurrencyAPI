package Config

import (
	"github.com/go-redis/redis/v8"
)

const SECRET = "TOP_SECRET"

var RDS *redis.Client

var RATES = map[string]interface{}{
	"TRY-USD":  0.054,
	"TRY-EURO": 0.051,
	"USD-TRY":  18.55,
	"USD-EURO": 0.95,
	"EURO-TRY": 19.63,
	"EURO-USD": 1.05,
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // use default DB
	})
	RDS = client

}

func GetRedisInstance() *redis.Client {
	return RDS
}
