package dto

import (
	"go-backend-template/internal/model"
)

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
