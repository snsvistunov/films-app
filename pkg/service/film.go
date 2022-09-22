package service

import (
	"github.com/snsvistunov/films-app/models"
	"github.com/snsvistunov/films-app/pkg/repository"
)

type FilmService struct {
	repo repository.Film
}

func NewFilmService(repo repository.Film) *FilmService {
	return &FilmService{repo: repo}
}

func (s *FilmService) Create(userID string, film models.Film) (string, error) {
	return s.repo.Create(userID, film)
}

func (s *FilmService) CheckFilmExists(name string) (bool, error) {
	return s.repo.CheckFilmExists(name)
}

func (s *FilmService) CheckDirectorExists(id string) (bool, error) {
	return s.repo.CheckDirectorExists(id)
}
