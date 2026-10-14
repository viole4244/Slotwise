package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis is the shared Redis client. Use it anywhere: internal.Redis
var Redis *redis.Client

// RedisConn connects to Redis and checks it is reachable.
func RedisConn(ctx context.Context, url string) error {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return fmt.Errorf("redis: parse url: %w", err)
	}

	client := redis.NewClient(opts)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		client.Close()
		return fmt.Errorf("redis: ping: %w", err)
	}

	Redis = client
	return nil
}
