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

	_, err := r.db.Exec(queryCreateEvent, e.UserID, e.Name, e.Description, e.AppointedDate)

	return err
}

func (r *EventRepository) GetAll(userID int64) ([]model.Event, error) {

	rows, err := r.db.Query(queryGetAllEvents, userID)
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

	err := r.db.QueryRow(queryGetEventByID, eventID).Scan(&e.Name, &e.Description, &e.AppointedDate)

	return e, err
}

func (r *EventRepository) Update(e *model.Event) error {

	_, err := r.db.Exec(queryUpdateEvent, e.Name, e.Description, e.AppointedDate, e.ID)
	return err
}

func (r *EventRepository) Delete(eventID int64) error {
	_, err := r.db.Exec(queryDeleteEvent, eventID)
	return err
}
