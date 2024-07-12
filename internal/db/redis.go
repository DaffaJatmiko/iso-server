package db

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

func InitRedis(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return rdb
}

func SetBlacklistToken(rdb *redis.Client, token string, duration time.Duration) error {
	ctx := context.Background()
	return rdb.Set(ctx, token, "blacklisted", duration).Err()
}

func IsBlacklistedToken(rdb *redis.Client, token string) (bool, error) {
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return val == "blacklisted", nil
}
