package postgres

import (
	"database/sql"
	"toDoApp/pkg/model"

	"github.com/sirupsen/logrus"
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
	query := "SELECT task.title, task.description, task.is_important, events.name, events.description FROM task INNER JOIN events ON task.user_id = events.user_id WHERE task.user_id = $1 AND events.user_id = $1 AND task.due_date = $2 AND events.appointed_date = $2"
	rows, err := r.db.Query(query, userID, date)
	if err != nil {
		logrus.Info(err)
		return nil, err
	}
	logrus.Info(rows)
	return nil, nil
}
