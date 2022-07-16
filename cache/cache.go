package cache

import (
	"context"
	"time"

	"bot/commands/pirate"

	"github.com/go-redis/redis/v9"
)

type PirateEntry struct {
	Metadata []pirate.Metadata
	Site     string
}

var rdb *redis.Client
var ctx = context.Background()

func Connect() {

	//redis://default:f7qdLKe6OvswaAmKQ6gLGOLJKKnoDPnL@redis-11314.c2.eu-west-1-3.ec2.cloud.redislabs.com:11314

	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-11314.c2.eu-west-1-3.ec2.cloud.redislabs.com:11314",
		Username: "default",
		Password: "f7qdLKe6OvswaAmKQ6gLGOLJKKnoDPnL",
	})
}

func Set(token string, value interface{}, ttl time.Duration) error {

	return rdb.Set(ctx, token, value, ttl).Err()
}

func Get(token string) (string, error) {

	return rdb.Get(ctx, token).Result()
}
