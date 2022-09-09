package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = `"user"`
	filmsTable      = "film"
	rolesTable      = "roles"
	userRolesTable  = "userRoles"
	directorsTable  = "director"
	favouritesTable = "favourites"
	wishlistTable   = "wishlist"
)

type PGConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg PGConfig) (*sqlx.DB, error) {
	dbType := "postgres"
	connectionParams := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open(dbType, connectionParams)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
