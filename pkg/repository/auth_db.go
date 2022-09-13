package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type AuthDB struct {
	db      *sqlx.DB
	storage *redis.Client
}

func NewAuthDB(db *sqlx.DB, storage *redis.Client) *AuthDB {
	return &AuthDB{
		db:      db,
		storage: storage,
	}
}
