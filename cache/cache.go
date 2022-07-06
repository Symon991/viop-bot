package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func Connect() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:49153",
		Username: "default",
		Password: "redispw",
	})
}

func Set(token string, value interface{}, ttl time.Duration) error {

	return rdb.Set(ctx, token, value, ttl).Err()
}

func Get(token string) (string, error) {

	return rdb.Get(ctx, token).Result()
}
