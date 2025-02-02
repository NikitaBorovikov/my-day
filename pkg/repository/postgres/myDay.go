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
	query := "SELECT task.title, task.description, task.is_important, events.name, events.description FROM task INNER JOIN events ON task.user_id = events.user_id WHERE task.user_id = $1 AND events.user_id = $1 AND task.due_date = $2 AND events.appointed_date = $2"
	rows, err := r.db.Query(query, userID, date)
	if err != nil {
		return nil, err
	}

	myDay := &model.MyDay{}

	for rows.Next() {
		task := model.Task{}
		event := model.Event{}

		if err := rows.Scan(
			&task.Title, &task.Description, &task.IsImportant, &event.Name, &event.Description); err != nil {
			continue
		}
		myDay.Tasks = append(myDay.Tasks, task)
		myDay.Events = append(myDay.Events, event)

	}

	return myDay, nil
}
