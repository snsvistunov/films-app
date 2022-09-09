package service

import (
	"github.com/snsvistunov/films-app/models"
	"github.com/snsvistunov/films-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CheckUserExist(user models.User) (bool, error)
}

type FilmsList interface {
}

type Film interface {
}

type Wishlist interface {
}

type Favourites interface {
}

type Service struct {
	Authorization
	FilmsList
	Film
	Wishlist
	Favourites
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
