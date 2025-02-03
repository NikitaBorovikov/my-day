package postgres

import (
	"toDoApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepositoty(db *sqlx.DB) model.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SignUp(u *model.User) error {
	_, err := r.db.Exec("INSERT INTO users (user_name, email, enc_password, reg_date) VALUES ($1, $2, $3, $4)",
		u.UserName, u.Email, u.EncPassword, u.RegDate)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) SignIn(email, password string) (*model.User, error) {
	u := &model.User{}

	if err := r.db.QueryRow("SELECT id, user_name, enc_password FROM users WHERE email = $1", email).Scan(
		&u.ID, &u.UserName, &u.EncPassword); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Delete(userID int64) error {

	errorChannel := make(chan error, 3)

	go func() {
		_, err := r.db.Exec("DELETE FROM users WHERE id = $1", userID)
		errorChannel <- err
	}()

	go func() {
		_, err := r.db.Exec("DELETE FROM task WHERE user_id = $1", userID)
		errorChannel <- err
	}()

	go func() {
		_, err := r.db.Exec("DELETE FROM events WHERE user_id = $1", userID)
		errorChannel <- err
	}()

	go func() {
		defer close(errorChannel)
		for i := 0; i < 3; i++ {
			<-errorChannel
		}
	}()

	firstError := findFirstError(errorChannel)
	return firstError
}

func findFirstError(errorChannel chan error) error {
	var firstError error
	for err := range errorChannel {
		if err != nil {
			firstError = err
		}
	}

	return firstError
}
