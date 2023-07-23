package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	*redis.Client
}

func NewRedisService(client *redis.Client) *Client {
	return &Client{
		Client: client,
	}
}

func (c *Client) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return c.Client.Set(ctx, key, value, ttl).Err()
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}

	return val, nil
}

