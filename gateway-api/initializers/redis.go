package initializers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() (*redis.Client, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}

	return rdb, nil
}

func CloseRedisConnection(rdb *redis.Client) error {
	ctx := context.Background()
	return rdb.Close()
}

func SetRedisKey(rdb *redis.Client, key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set redis key: %v", err)
	}
	return nil
}

func GetRedisKey(rdb *redis.Client, key string) (string, error) {
	ctx := context.Background()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get redis key: %v", err)
	}
	return val, nil
}
