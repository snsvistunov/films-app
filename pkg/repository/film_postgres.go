package repository

import (
	"database/sql"
	"fmt"

	"github.com/snsvistunov/films-app/models"
)

func (r *FilmPostgres) Create(userID string, film models.Film) (string, error) {
	var filmID string

	createFilmQuery := fmt.Sprintf("INSERT INTO %s (name, genre, director_id, rate, year, minutes) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id", filmsTable)
	row := r.db.QueryRow(createFilmQuery, film.Name, film.Genre, film.DirectorId, film.Rate, film.Year, film.Minutes)
	if err := row.Scan(&filmID); err != nil {
		return "", err
	}

	return filmID, nil
}

func (r *FilmPostgres) CheckFilmExists(name string) (bool, error) {
	var existedFilmName string

	query := fmt.Sprintf("SELECT name FROM %s WHERE name = '%s'", filmsTable, name)
	row := r.db.QueryRow(query)
	if err := row.Scan(&existedFilmName); err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if existedFilmName == name {
		return true, nil
	}

	return false, nil
}

func (r *FilmPostgres) CheckDirectorExists(id string) (bool, error) {
	var existedDirectorID string

	query := fmt.Sprintf("SELECT id FROM %s WHERE id = '%s'", directorsTable, id)
	fmt.Println("directorID", id)            //don't forget to del
	fmt.Println("query to director ", query) //don't forget to del
	row := r.db.QueryRow(query)
	if err := row.Scan(&existedDirectorID); err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if existedDirectorID == id {
		return true, nil
	}

	return false, nil
}
