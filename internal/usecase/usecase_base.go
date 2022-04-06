package usecase

import (
	"time"

	"go-backend-template/internal/database"
)

func NewUsecases(db database.Service, config Config) Usecases {
	return Usecases{
		Auth: &AuthUsecases{db: db, config: config},
		User: &UserUsecases{db: db, config: config},
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
