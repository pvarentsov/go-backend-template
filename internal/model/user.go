package model

import (
	"go-backend-template/internal/util/crypto"
	"go-backend-template/internal/util/errors"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func NewUser(firstName, lastName, email, password string) (User, error) {
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
	if err := user.Validate(); err != nil {
		return User{}, err
	}
	if err := user.HashPassword(); err != nil {
		return User{}, err
	}

	return user, nil
}

func (user *User) UpdateInfo(firstName, lastName, email string) error {
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

func (user *User) ChangePassword(newPassword string) error {
	user.Password = newPassword

	if err := user.HashPassword(); err != nil {
		return err
	}

	return user.Validate()
}

func (user *User) Validate() error {
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

func (user *User) ComparePassword(password string) bool {
	return crypto.CompareHashAndPassword(user.Password, password)
}

func (user *User) HashPassword() error {
	hash, err := crypto.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	return nil
}
