package usecases

import (
	"toDoApp/pkg/repository"
)

type UseCases struct {
	UserUseCase  *UserUseCase
	TaskUseCase  *TaskUseCase
	EventUseCase *EventUseCase
	MyDayUseCase *MyDayUseCase
}

func InitUseCases(r *repository.Repository) *UseCases {
	return &UseCases{
		UserUseCase:  NewUserUseCase(r.UserRepository),
		TaskUseCase:  NewTaskUseCase(r.TaskRepository),
		EventUseCase: NewEventUseCase(r.EventRepository),
		MyDayUseCase: NewMyDayUseCase(r.MyDayRepository),
	}
}
