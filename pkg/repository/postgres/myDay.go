package postgres

import (
	"toDoApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type MyDayRepository struct {
	db *sqlx.DB
}

func NewMyDayRepository(db *sqlx.DB) model.MyDayRepository {
	return &MyDayRepository{
		db: db,
	}
}

func (r *MyDayRepository) Get(userID int64, date string) (*model.MyDay, error) {
	rows, err := r.db.Query(queryGetMyDay, userID, date)
	if err != nil {
		return nil, err
	}

	myDay := &model.MyDay{}

	for rows.Next() {
		task := model.Task{}
		event := model.Event{}

		if err := rows.Scan(
			&task.Title, &task.Description, &task.IsImportant, &task.IsDone, &event.Name, &event.Description); err != nil {
			continue
		}
		myDay.Tasks = append(myDay.Tasks, task)
		myDay.Events = append(myDay.Events, event)

	}

	return myDay, nil
}
