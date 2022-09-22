package repository

import (
	"database/sql"
	"fmt"

	"github.com/snsvistunov/films-app/models"
)

func (r *AuthDB) CreateUser(user models.User) (string, error) {
	var userID string
	var roleID byte

	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("INSERT INTO %s (login, password, age) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Age)

	if err := row.Scan(&userID); err != nil {
		if err := tx.Rollback(); err != nil {
			return "", err
		}
		return "", err
	}

	query = fmt.Sprintf("SELECT id FROM %s WHERE name=$1", rolesTable)
	row = r.db.QueryRow(query, baseUserRole)

	if err := row.Scan(&roleID); err != nil {
		if err := tx.Rollback(); err != nil {
			return "", err
		}
		return "", err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, roles_id) values ($1, $2)", userRolesTable)
	if _, err := r.db.Query(query, userID, roleID); err != nil {
		if err := tx.Rollback(); err != nil {
			return "", err
		}
		return "", err
	}

	return userID, tx.Commit()
}

func (r *AuthDB) CheckUserExists(login string) (bool, error) {
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

func (r *AuthDB) GetUser(login string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1", usersTable)
	err := r.db.Get(&user, query, login)

	return user, err
}

func (r *AuthDB) GetUserRole(userID string) (string, error) {
	var userRole string
	query := fmt.Sprintf("SELECT name FROM %s WHERE id IN (SELECT roles_id FROM %s WHERE user_id = $1)", rolesTable, userRolesTable)
	err := r.db.Get(&userRole, query, userID)

	return userRole, err
}
