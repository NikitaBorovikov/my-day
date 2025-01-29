package db

import (
	"database/sql"
	"fmt"
	"toDoApp/pkg/config"
)

func Connect(config *config.Config) (*sql.DB, error) {
	dbConnectionString := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User,
		config.Postgres.Password, config.Postgres.Name)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
