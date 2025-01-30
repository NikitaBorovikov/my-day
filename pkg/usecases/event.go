package usecases

import (
	"time"
	"toDoApp/pkg/model"

	"github.com/go-playground/validator"
)

type EventUseCase struct {
	EventRepository model.EventRepository
}

func NewEventUseCase(eventRepository model.EventRepository) *EventUseCase {
	return &EventUseCase{
		EventRepository: eventRepository,
	}
}

func (uc *EventUseCase) Create(e *model.Event) error {

	if err := validateForEvent(e); err != nil {
		return err
	}

	if err := setDateFormatForEvent(e); err != nil {
		return err
	}

	if err := uc.EventRepository.Create(e); err != nil {
		return err
	}

	return nil
}

func (uc *EventUseCase) GetAll(userID int64) ([]model.Event, error) {
	events, err := uc.EventRepository.GetAll(userID)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (uc *EventUseCase) GetByID(eventID int64) (*model.Event, error) {
	event, err := uc.EventRepository.GetByID(eventID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (uc *EventUseCase) Update(e *model.Event) error {
	return nil
}

func (uc *EventUseCase) Delete(eventID int64) error {
	err := uc.EventRepository.Delete(eventID)
	return err

}

func validateForEvent(e *model.Event) error {
	validate := validator.New()

	if err := validate.Struct(e); err != nil {
		return err
	}

	return nil
}

func setDateFormatForEvent(e *model.Event) error {
	if e.AppointedDate == "" {
		e.AppointedDate = time.Date(1970, 01, 01, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
		return nil
	}

	parsedDate, err := time.Parse("01-02-2006", e.AppointedDate)
	if err != nil {
		return err
	}

	e.AppointedDate = parsedDate.Format(time.RFC3339)
	return nil
}
