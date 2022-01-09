package usecase

import (
	"go-backend-template/internal/database"
	"time"
)

func NewUsecases(dbService database.Service, config Config) Usecases {
	return Usecases{
		Auth: &AuthUsecases{dbService: dbService, config: config},
		User: &UserUsecases{dbService: dbService, config: config},
	}
}

type Usecases struct {
	Auth *AuthUsecases
	User *UserUsecases
}

type Config interface {
	AccessTokenSecret() string
	AccessTokenExpiresDate() time.Time
}
