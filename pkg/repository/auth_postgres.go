package repository

import (
	"database/sql"
	"fmt"

	"github.com/snsvistunov/films-app/models"
)

func (r *AuthDB) CreateUser(user models.User) (string, error) {
	var userId string
	var roleId byte

	query := fmt.Sprintf("INSERT INTO %s (login, password, age) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Age)

	if err := row.Scan(&userId); err != nil {
		return "", err
	}

	baseUserRole := "user"
	query = fmt.Sprintf("SELECT id FROM %s WHERE name=$1", rolesTable)
	row = r.db.QueryRow(query, baseUserRole)

	if err := row.Scan(&roleId); err != nil {
		return "", err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, roles_id) values ($1, $2)", userRolesTable)
	if _, err := r.db.Query(query, userId, roleId); err != nil {
		return "", err
	}

	return userId, nil
}

func (r *AuthDB) CheckUserExist(login string) (bool, error) {
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
