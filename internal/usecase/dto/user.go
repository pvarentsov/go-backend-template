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

func (dto User) MapFrom(user model.User) User {
	dto.Id = user.Id
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email

	return dto
}

// UserUpdateInfo

type UserUpdateInfo struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// UserLogin

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserAdd

type UserAdd struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (dto *UserAdd) MapTo() (model.User, error) {
	return model.NewUser(
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
	)
}

// UserLoggedInfo

type UserLoggedInfo struct {
	User
	Token string `json:"token"`
}

func (dto UserLoggedInfo) MapFrom(user model.User, token string) UserLoggedInfo {
	dto.Id = user.Id
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email
	dto.Token = token

	return dto
}

// UserChangePassword

type UserChangePassword struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}
