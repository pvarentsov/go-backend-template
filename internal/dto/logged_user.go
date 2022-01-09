package dto

import "go-backend-template/internal/model"

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
