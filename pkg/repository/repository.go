package repository

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/snsvistunov/films-app/models"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CheckUserExist(login string) (bool, error)
	GetUser(login string) (models.User, error)
	SaveToken(userID []uint8, token string, ttl time.Duration) error
	DeleteToken(token string) error
	GetUserID(token string) (string, error)
}

type FilmsList interface {
}

type Film interface {
}

type Wishlist interface {
}

type Favourites interface {
}

type Repository struct {
	Authorization
	FilmsList
	Film
	Wishlist
	Favourites
}

func NewRepository(db *sqlx.DB, storage *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db, storage),
	}
}
