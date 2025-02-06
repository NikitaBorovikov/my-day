package postgres

import (
	"toDoApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) model.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) Create(e *model.Event) error {

	query := "INSERT INTO events (user_id, name, description, appointed_date) VALUES ($1, $2, $3, $4)"

	_, err := r.db.Exec(query, e.UserID, e.Name, e.Description, e.AppointedDate)

	return err
}

func (r *EventRepository) GetAll(userID int64) ([]model.Event, error) {

	query := "SELECT name, description, appointed_date FROM events WHERE user_id = $1"

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	allEvents := []model.Event{}
	for rows.Next() {
		e := model.Event{}
		err := rows.Scan(&e.Name, &e.Description, &e.AppointedDate)
		if err != nil {
			continue
		}
		allEvents = append(allEvents, e)
	}
	return allEvents, nil
}

func (r *EventRepository) GetByID(eventID int64) (*model.Event, error) {
	e := &model.Event{}

	query := "SELECT name, description, appointed_date FROM events WHERE id = $1"

	err := r.db.QueryRow(query, eventID).Scan(&e.Name, &e.Description, &e.AppointedDate)

	return e, err
}

func (r *EventRepository) Update(e *model.Event) error {

	query := "UPDATE events SET name = $1, description = $2, appointed_date = $3 WHERE id = $4"

	_, err := r.db.Exec(query, e.Name, e.Description, e.AppointedDate, e.ID)
	return err
}

func (r *EventRepository) Delete(eventID int64) error {
	query := "DELETE FROM events WHERE id = $1"
	_, err := r.db.Exec(query, eventID)
	return err
}
