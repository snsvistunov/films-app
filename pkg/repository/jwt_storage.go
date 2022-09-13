package repository

import (
	"github.com/go-redis/redis"
)

type AuthRedis struct {
	storage *redis.Client
}

func NewAuthRedis(storage *redis.Client) *AuthRedis {
	return &AuthRedis{
		storage: storage,
	}
}
