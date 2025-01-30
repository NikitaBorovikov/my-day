package usecases

import (
	"time"
	"toDoApp/pkg/model"

	"github.com/go-playground/validator"
)

type TaskUseCase struct {
	TaskRepository model.TaskRepository
}

func NewTaskUseCase(taskRepository model.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (uc *TaskUseCase) Create(t *model.Task) error {

	if err := validateForTask(t); err != nil {
		return err
	}

	if err := setDateFormatForTask(t); err != nil {
		return err
	}

	if err := uc.TaskRepository.Create(t); err != nil {
		return err
	}

	return nil
}

func (uc *TaskUseCase) GetAll(userID int64) ([]model.Task, error) {

	allTasks, err := uc.TaskRepository.GetAll(userID)
	if err != nil {
		return nil, err
	}
	return allTasks, nil
}

func (uc *TaskUseCase) GetByID(taskID int64) (*model.Task, error) {
	task, err := uc.TaskRepository.GetByID(taskID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (uc *TaskUseCase) Update(t *model.Task) error {
	if err := validateForTask(t); err != nil {
		return err
	}

	if err := setDateFormatForTask(t); err != nil {
		return err
	}

	err := uc.TaskRepository.Update(t)
	return err

}

func (uc *TaskUseCase) Delete(taskID int64) error {
	err := uc.TaskRepository.Delete(taskID)
	return err
}

func validateForTask(t *model.Task) error {
	validate := validator.New()
	err := validate.Struct(t)
	return err
}

func setDateFormatForTask(t *model.Task) (err error) {
	t.DueDate, err = setDueDate(t.DueDate)
	t.CreatedDate = time.Now().Format(time.RFC3339)
	return err
}

func setDueDate(dueDate string) (string, error) {
	if dueDate == "" {
		formattedDate := time.Date(1970, 01, 01, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
		return formattedDate, nil
	}

	parsedDate, err := time.Parse("01-02-2006", dueDate)
	if err != nil {
		return "", err
	}
	formattedDate := parsedDate.Format(time.RFC3339)
	return formattedDate, nil

}
