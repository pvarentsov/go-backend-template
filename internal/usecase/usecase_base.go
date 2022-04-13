package usecase

import (
	"time"

	"go-backend-template/internal/database"
	"go-backend-template/internal/util/crypto"
)

func NewUsecases(db database.Service, config Config, crypto crypto.Crypto) Usecases {
	return Usecases{
		Auth: &authUsecases{db: db, crypto: crypto, config: config},
		User: &userUsecases{db: db, crypto: crypto, config: config},
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
