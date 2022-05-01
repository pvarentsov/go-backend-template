package user

import (
	"go-backend-template/internal/base/crypto"
	"go-backend-template/internal/base/errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserModel struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func NewUser(firstName, lastName, email, password string) (UserModel, error) {
	user := UserModel{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
	if err := user.Validate(); err != nil {
		return UserModel{}, err
	}

	return user, nil
}

func (user *UserModel) Update(firstName, lastName, email string) error {
	if len(firstName) > 0 {
		user.FirstName = firstName
	}
	if len(lastName) > 0 {
		user.LastName = lastName
	}
	if len(email) > 0 {
		user.Email = email
	}

	return user.Validate()
}

func (user *UserModel) ChangePassword(newPassword string, crypto crypto.Crypto) error {
	user.Password = newPassword

	if err := user.HashPassword(crypto); err != nil {
		return err
	}

	return user.Validate()
}

func (user *UserModel) Validate() error {
	err := validation.ValidateStruct(user,
		validation.Field(&user.FirstName, validation.Required, validation.Length(2, 100)),
		validation.Field(&user.LastName, validation.Required, validation.Length(2, 100)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(5, 100)),
	)
	if err != nil {
		return errors.New(errors.ValidationError, err.Error())
	}

	return nil
}

func (user *UserModel) ComparePassword(password string, crypto crypto.Crypto) bool {
	return crypto.CompareHashAndPassword(user.Password, password)
}

func (user *UserModel) HashPassword(crypto crypto.Crypto) error {
	hash, err := crypto.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	return nil
}
