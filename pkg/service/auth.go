package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/snsvistunov/films-app/models"
	"github.com/snsvistunov/films-app/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenTTL     = 24 * time.Hour
	signinKey    = "jh23jerJH23ndlq2#k19Jn1nc1&"
	signinMethod = "HS256"
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

func (s *AuthService) CheckUserExists(login string) (bool, error) {
	return s.repo.CheckUserExists(login)
}

func (s *AuthService) GetUserID(token string) (string, error) {
	return s.repo.GetUserID(token)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	token := jwt.New(jwt.GetSigningMethod(signinMethod))

	token.Claims = &JWTClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	}

	tokenSignedString, err := token.SignedString([]byte(signinKey))
	if err != nil {
		return "", err
	}

	if err := s.SaveToken(user.Id, tokenSignedString, tokenTTL); err != nil {
		return "", err
	}

	return token.SignedString([]byte(signinKey))
}

func (s *AuthService) SaveToken(userID []uint8, token string, ttl time.Duration) error {
	return s.repo.SaveToken(userID, token, ttl)
}

func (s *AuthService) DeleteToken(token string) error {
	return s.repo.DeleteToken(token)
}

func (s *AuthService) GetUserRole(userID string) (string, error) {
	return s.repo.GetUserRole(userID)
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *AuthService) IsAdmin(userID string, admin string) (bool, error) {
	role, err := s.GetUserRole(userID)
	if err != nil {
		return false, err
	}

	if role != admin {
		return false, nil
	}
	return true, nil
}
