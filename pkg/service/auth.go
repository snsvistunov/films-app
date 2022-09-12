package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/snsvistunov/films-app/models"
	"github.com/snsvistunov/films-app/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenTTL  = 24 * time.Hour
	signinKey = "jh23jerJH23ndlq2#k19Jn1nc1&"
)

type JWTClaims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

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

func (s *AuthService) CheckUserExist(login string) (bool, error) {
	return s.repo.CheckUserExist(login)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = &JWTClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	}

	return token.SignedString([]byte(signinKey))
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
