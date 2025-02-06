package usecases

import (
	"errors"
	"time"
	"toDoApp/pkg/model"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository model.UserRepository
}

func NewUserUseCase(userRepository model.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *UserUseCase) SignUp(u *model.User) error {

	setDefaultValuesForUser(u)

	if err := validateForSignUp(u); err != nil {
		return err
	}

	if err := encryptPassword(u); err != nil {
		return err
	}

	if err := uc.UserRepository.SignUp(u); err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) SignIn(email, password string) (*model.User, error) {
	u, err := uc.UserRepository.SignIn(email, password)
	if err != nil {
		return nil, err
	}
	if !comparePassword(password, u.EncPassword) {
		return nil, errors.New("incorrect password")
	}

	return u, nil
}

func (uc *UserUseCase) Get(userID int64) (*model.User, error) {
	user, err := uc.UserRepository.Get(userID)
	return user, err
}

func (uc *UserUseCase) Delete(userID int64) error {
	err := uc.UserRepository.Delete(userID)
	return err
}

func setDefaultValuesForUser(u *model.User) {
	u.RegDate = time.Now().Format(time.RFC3339)
}

func validateForSignUp(u *model.User) error {
	validate := validator.New()
	err := validate.Struct(u)
	return err
}

func encryptPassword(u *model.User) error {
	enc_password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.EncPassword = string(enc_password)
	return nil
}

func comparePassword(enteredPassword, realHashPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(realHashPassword), []byte(enteredPassword)); err != nil {
		return false
	}
	return true
}
