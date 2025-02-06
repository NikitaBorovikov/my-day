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

	_, err := r.db.Exec(queryCreateTask, t.UserID, t.Title, t.Description, t.IsImportant, t.DueDate, t.CreatedDate, t.IsDone)
	return err
}

func (r *TaskRepository) GetAll(userID int64) ([]model.Task, error) {

	rows, err := r.db.Query(queryGetAllTask, userID)
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

	err := r.db.QueryRow(queryGetTaskByID, taskID).Scan(
		&task.UserID, &task.Title, &task.Description, &task.IsImportant, &task.IsDone, &task.DueDate, &task.CreatedDate)

	return task, err
}

func (r *TaskRepository) Update(t *model.Task) error {

	_, err := r.db.Exec(queryUpdateTask, t.Title, t.Description, t.IsImportant, t.DueDate, t.IsDone, &t.ID)
	return err
}

func (r *TaskRepository) Delete(taskID int64) error {

	_, err := r.db.Exec(queryDeleteTask, taskID)
	return err
}
