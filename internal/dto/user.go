package dto

import (
	"go-backend-template/internal/model"
)

// User

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (dto *User) MapFrom(user model.User) {
	dto.Id = user.Id
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email
}

// UpdateUserInfo

type UpdateUserInfo struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// LoginUser

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AddUser

type AddUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (dto *AddUser) MapTo() (model.User, error) {
	return model.NewUser(
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
	)
}

// LoggedUser

type LoggedUser struct {
	User
	Token string `json:"token"`
}

func (dto *LoggedUser) MapFrom(user model.User, token string) {
	dto.Id = user.Id
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email
	dto.Token = token
}

// ChangeUserPassword

type ChangeUserPassword struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}
