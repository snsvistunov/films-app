package service

import (
	"github.com/snsvistunov/films-app/models"
	"github.com/snsvistunov/films-app/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (string, error) {

	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hash
	return s.repo.CreateUser(user)
}

func (s *AuthService) CheckUserExist(user models.User) (bool, error) {
	return s.repo.CheckUserExist(user)
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
