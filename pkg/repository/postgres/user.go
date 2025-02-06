package postgres

import (
	"toDoApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) model.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SignUp(u *model.User) error {

	_, err := r.db.Exec(querySignUp, u.UserName, u.Email, u.EncPassword, u.RegDate)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) SignIn(email, password string) (*model.User, error) {
	u := &model.User{}

	err := r.db.QueryRow(querySignIn, email).Scan(&u.ID, &u.UserName, &u.EncPassword)
	return u, err
}

func (r *UserRepository) Get(userID int64) (*model.User, error) {
	u := &model.User{}

	err := r.db.QueryRow(queryGetProfile, userID).Scan(&u.UserName, &u.Email, &u.RegDate)
	return u, err
}

func (r *UserRepository) Delete(userID int64) error {

	errorChannel := make(chan error, 3)

	go func() {
		query := "DELETE FROM users WHERE id = $1"
		_, err := r.db.Exec(query, userID)
		errorChannel <- err
	}()

	go func() {
		query := "DELETE FROM task WHERE user_id = $1"
		_, err := r.db.Exec(query, userID)
		errorChannel <- err
	}()

	go func() {
		query := "DELETE FROM events WHERE user_id = $1"
		_, err := r.db.Exec(query, userID)
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
