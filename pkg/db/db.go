package db

import (
	"fmt"
	"toDoApp/pkg/config"

	"github.com/jmoiron/sqlx"
)

func Connect(config *config.Config) (*sqlx.DB, error) {
	dbConnectionString := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User,
		config.Postgres.Password, config.Postgres.Name)

	db, err := sqlx.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
