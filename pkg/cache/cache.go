// pkg/cache/cache.go
// Used by: Auth, Health, Tenant, Repo, API Gateway
package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect(url string) {
	Client = redis.NewClient(&redis.Options{Addr: url})
}

func Set(ctx context.Context, key string, val interface{}) error {
	return Client.Set(ctx, key, val, 0).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	return Client.Get(ctx, key).Result()
}
