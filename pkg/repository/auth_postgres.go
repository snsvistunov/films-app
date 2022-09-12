package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/snsvistunov/films-app/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user models.User) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (login, password, age) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Age)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *AuthPostgres) CheckUserExist(login string) (bool, error) {
	var existedLogin string

	query := fmt.Sprintf("SELECT login FROM %s WHERE login = '%s'", usersTable, login)
	row := r.db.QueryRow(query)
	if err := row.Scan(&existedLogin); err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if existedLogin == login {
		return true, nil
	}

	return false, nil
}

func (r *AuthPostgres) GetUser(login string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1", usersTable)
	err := r.db.Get(&user, query, login)
	fmt.Printf("\nUser %v\n", user.Password)
	return user, err
}
