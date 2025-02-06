package postgres

import (
	"toDoApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) model.TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(t *model.Task) error {

	query := "INSERT INTO task (user_id, title, description, is_important, due_date, created_date, is_done) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := r.db.Exec(query, t.UserID, t.Title, t.Description, t.IsImportant, t.DueDate, t.CreatedDate, t.IsDone)
	return err
}

func (r *TaskRepository) GetAll(userID int64) ([]model.Task, error) {

	query := "SELECT title, description, is_important, is_done, due_date, created_date FROM task WHERE user_id = $1"

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	allTasks := []model.Task{}

	for rows.Next() {
		t := model.Task{}
		if err := rows.Scan(
			&t.Title, &t.Description, &t.IsImportant, &t.IsDone, &t.DueDate, &t.CreatedDate); err != nil {
			continue
		}

		allTasks = append(allTasks, t)
	}
	return allTasks, nil
}

func (r *TaskRepository) GetByID(taskID int64) (*model.Task, error) {
	task := &model.Task{}

	query := "SELECT user_id, title, description, is_important, is_done, due_date, created_date FROM task WHERE id = $1"

	err := r.db.QueryRow(query, taskID).Scan(
		&task.UserID, &task.Title, &task.Description, &task.IsImportant, &task.IsDone, &task.DueDate, &task.CreatedDate)

	return task, err
}

func (r *TaskRepository) Update(t *model.Task) error {

	query := "UPDATE task SET title = $1, description = $2, is_important = $3, due_date = $4, is_done = $5 WHERE id = $6"

	_, err := r.db.Exec(query, t.Title, t.Description, t.IsImportant, t.DueDate, t.IsDone, &t.ID)
	return err
}

func (r *TaskRepository) Delete(taskID int64) error {

	query := "DELETE FROM task WHERE id = $1"

	_, err := r.db.Exec(query, taskID)
	return err
}
