package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var rdb *redis.Client = nil
var ctx = context.Background()

func GetRedis() (*redis.Client, context.Context) {
	if rdb != nil {
		return rdb, ctx
	}

	addr := viper.GetString("RDB_ADDR")
	password := viper.GetString("RDB_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis Init Failed")
	}
	return rdb, ctx
}

func SetValue(key, value string, time time.Duration) error {
	rdb, ctx := GetRedis()
	err := rdb.Set(ctx, key, value, time).Err()
	return err
}

func GetValue(key string) (string, error) {
	rdb, ctx := GetRedis()
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func DeleteValue(key string) error {
	rdb, ctx := GetRedis()
	err := rdb.Del(ctx, key).Err()
	if err == redis.Nil {
		return nil
	}
	return err
}
