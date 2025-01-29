package postgres

import (
	"database/sql"
	"toDoApp/pkg/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepositoty(db *sql.DB) model.UserRepository {
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
