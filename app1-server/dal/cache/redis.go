package redis

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "10.97.25.112:6379",
		Password: "123456",
		DB:       0,
	})
}

func GetCntAndInc(key string) (int, error) {
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		client.Set(ctx, key, "1", 0).Err()
		return 1, err
	}
	client.Incr(ctx, key).Err()
	cnt, _ := strconv.Atoi(val)
	return cnt, nil
}
