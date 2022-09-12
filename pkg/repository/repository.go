package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/snsvistunov/films-app/models"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CheckUserExist(login string) (bool, error)
	GetUser(login string) (models.User, error)
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
