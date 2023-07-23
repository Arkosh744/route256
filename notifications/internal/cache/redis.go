package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	r "route256/libs/client/redis"
	"route256/notifications/internal/models"
	"time"
)

type RedisCache struct {
	rc *r.Client
}

func NewRedis(client *r.Client) *RedisCache {
	return &RedisCache{rc: client}
}

func (c *RedisCache) GetUserHistoryDay(ctx context.Context, userID int64) ([]models.OrderMessage, error) {
	key := BuildUserIDHistoryKey(userID)

	data, err := c.rc.SMembers(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var messages []models.OrderMessage
	for _, item := range data {
		var msg models.OrderMessage
		err = json.Unmarshal([]byte(item), &msg)
		if err != nil {
			return nil, err
		}

		if msg.CreatedAt.Before(time.Now().Add(-24 * time.Hour)) {
			c.rc.SRem(ctx, key, item)
			continue
		}

		messages = append(messages, msg)

		orderKey := BuildGetUserIDByOrderKey(msg.OrderID)
		err = c.rc.Set(ctx, orderKey, userID, 24*time.Hour)
		if err != nil {
			return nil, err
		}

		latestKey := BuildLatestMsgTimeKey(msg.UserID)
		latestTime, err := c.rc.Get(ctx, latestKey)
		if err != nil {
			return nil, err
		}

		if latestTime == "" {
			if err = c.rc.Set(ctx, latestKey, msg.CreatedAt, 24*time.Hour); err != nil {
				return nil, err
			}
			continue
		}

		parsedTime, err := time.Parse(time.RFC3339, latestTime)
		if err != nil {
			return nil, err
		}

		if msg.CreatedAt.Unix() > parsedTime.Unix() {
			if err = c.rc.Set(ctx, latestKey, msg.CreatedAt, 24*time.Hour); err != nil {
				return nil, err
			}
		}

	}

	return messages, nil
}

func (c *RedisCache) AddToUserHistoryDay(ctx context.Context, msg models.OrderMessage) error {
	key := BuildUserIDHistoryKey(msg.UserID)

	if msg.CreatedAt.Before(time.Now().Add(-24 * time.Hour)) {
		return nil
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	cmd := c.rc.SAdd(ctx, key, jsonData)
	if err := cmd.Err(); err != nil {
		return err
	}

	orderKey := BuildGetUserIDByOrderKey(msg.OrderID)
	err = c.rc.Set(ctx, orderKey, msg.UserID, 24*time.Hour)
	if err != nil {
		return err
	}

	latestKey := BuildLatestMsgTimeKey(msg.UserID)
	latestTime, err := c.rc.Get(ctx, latestKey)
	if err != nil {
		return err
	}

	if latestTime == "" {
		if err = c.rc.Set(ctx, latestKey, msg.CreatedAt, 24*time.Hour); err != nil {
			return err
		}

		return nil
	}

	parsedTime, err := time.Parse(time.RFC3339, latestTime)
	if err != nil {
		return err
	}

	if msg.CreatedAt.Unix() > parsedTime.Unix() {
		if err = c.rc.Set(ctx, latestKey, msg.CreatedAt, 24*time.Hour); err != nil {
			return err
		}
	}

	return nil
}

func (c *RedisCache) GetLatestMessageTime(ctx context.Context, userID int64) (time.Time, error) {
	latestKey := BuildLatestMsgTimeKey(userID)

	latestTime, err := c.rc.Get(ctx, latestKey)
	if err != nil {
		return time.Time{}, err
	}

	if latestTime == "" {
		return time.Time{}, nil
	}

	parsedTime, err := time.Parse(time.RFC3339, latestTime)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return c.rc.Set(ctx, key, value, ttl)
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return c.rc.Get(ctx, key)
}
