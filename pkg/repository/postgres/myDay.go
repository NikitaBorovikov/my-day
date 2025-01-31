package postgres

import (
	"database/sql"
	"toDoApp/pkg/model"
)

type MyDayRepository struct {
	db *sql.DB
}

func NewMyDayRepository(db *sql.DB) model.MyDayRepository {
	return &MyDayRepository{
		db: db,
	}
}

func (r *MyDayRepository) Get(userID int64, date string) (*model.MyDay, error) {
	return nil, nil
}
