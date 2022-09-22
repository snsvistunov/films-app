package service

import (
	"time"

	"github.com/snsvistunov/films-app/models"
	"github.com/snsvistunov/films-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CheckUserExists(login string) (bool, error)
	GenerateToken(login, password string) (string, error)
	SaveToken(userID []uint8, token string, ttl time.Duration) error
	DeleteToken(token string) error
	GetUserID(token string) (string, error)
	GetUserRole(userID string) (string, error)
	IsAdmin(userID string, admin string) (bool, error)
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

type Service struct {
	Authorization
	Film
	Wishlist
	Favourites
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Film:          NewFilmService(repos.Film),
	}
}
