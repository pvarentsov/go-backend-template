package usecase

import (
	"time"

	"go-backend-template/internal/database"
)

func NewUsecases(db database.Service, config Config) Usecases {
	return Usecases{
		Auth: &authUsecases{db: db, config: config},
		User: &userUsecases{db: db, config: config},
	}
}

type Usecases struct {
	Auth AuthUsecases
	User UserUsecases
}

type Config interface {
	AccessTokenSecret() string
	AccessTokenExpiresDate() time.Time
}
