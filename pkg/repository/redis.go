package repository

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host string
	Port string
}

func NewRedisStorage(cfg RedisConfig) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	storage := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := storage.Ping().Result()
	if err != nil {
		return nil, err
	}

	return storage, nil
}
