package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/symon991/pirate/sites"
)

type PirateEntry struct {
	Metadata []sites.Metadata
	Site     string
}

var rdb *redis.Client
var ctx = context.Background()

func Connect() {

	//redis://default:nOkwJjqxqrkAzQJSX12eVquZ8IpA9htw@redis-13631.c6.eu-west-1-1.ec2.cloud.redislabs.com:13631

	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-13631.c6.eu-west-1-1.ec2.cloud.redislabs.com:13631",
		Username: "default",
		Password: "nOkwJjqxqrkAzQJSX12eVquZ8IpA9htw",
	})
}

func Set(token string, value interface{}, ttl time.Duration) error {

	return rdb.Set(ctx, token, value, ttl).Err()
}

func Get(token string) (string, error) {

	return rdb.Get(ctx, token).Result()
}
