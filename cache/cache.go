package cache

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func Connect(url *url.URL) {

	addr := fmt.Sprintf("%s:%s", url.Hostname(), url.Port())
	username := url.User.Username()
	password, _ := url.User.Password()

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
	})
}

func Set(token string, value interface{}, ttl time.Duration) error {

	return rdb.Set(ctx, token, value, ttl).Err()
}

func Get(token string) (string, error) {

	return rdb.Get(ctx, token).Result()
}
