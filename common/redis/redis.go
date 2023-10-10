package redisstore

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisStore struct {
	*redis.Client
}

func NewRedis() *RedisStore {
	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: "",
		DB:       0,
	})
	return &RedisStore{r}
}

func (r *RedisStore) Set(key string, value interface{}, t time.Duration) error {
	return r.Client.Set(context.Background(), key, value, t).Err()
}

func (r *RedisStore) Get(key string) (string, error) {
	val, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisStore) Exists(key string) bool {
	_, err := r.Client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return false
	}
	return true
}

func (r *RedisStore) Delete(key string) (int64, error) {
	deleted, err := r.Client.Del(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
