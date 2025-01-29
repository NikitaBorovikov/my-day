package postgres

import (
	"database/sql"
	"toDoApp/pkg/model"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) model.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) Create(e *model.Event) error {
	_, err := r.db.Exec("INSERT INTO events (user_id, name, description, appointed_date) VALUES ($1, $2, $3, $4)",
		e.UserID, e.Name, e.Description, e.AppointedDate)

	return err
}

func (r *EventRepository) GetAll(userID int64) ([]model.Event, error) {
	return nil, nil
}

func (r *EventRepository) GetByID(eventID int64) (*model.Event, error) {
	return nil, nil
}

func (r *EventRepository) Update(e *model.Event) error {
	return nil
}

func (r *EventRepository) Delete(eventID int64) error {
	return nil
}
