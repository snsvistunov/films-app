package repository

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/snsvistunov/films-app/models"
)

var baseUserRole = "user"

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CheckUserExists(login string) (bool, error)
	GetUser(login string) (models.User, error)
	SaveToken(userID []uint8, token string, ttl time.Duration) error
	DeleteToken(token string) error
	GetUserID(token string) (string, error)
	GetUserRole(userID string) (string, error)
}

type Film interface {
	Create(userID string, film models.Film) (string, error)
	CheckFilmExists(name string) (bool, error)
	CheckDirectorExists(id string) (bool, error)
}

type Wishlist interface {
}

type Favourites interface {
}

type Repository struct {
	Authorization
	Film
	Wishlist
	Favourites
}

func NewRepository(db *sqlx.DB, storage *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db, storage),
		Film:          NewFilmPostgres(db),
	}
}
