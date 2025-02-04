package repository

import (
	"toDoApp/pkg/model"
	"toDoApp/pkg/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepository  model.UserRepository
	TaskRepository  model.TaskRepository
	EventRepository model.EventRepository
	MyDayRepository model.MyDayRepository
}

func InitRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:  postgres.NewUserRepository(db),
		TaskRepository:  postgres.NewTaskRepository(db),
		EventRepository: postgres.NewEventRepository(db),
		MyDayRepository: postgres.NewMyDayRepository(db),
	}
}
