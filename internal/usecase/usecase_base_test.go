package usecase

import (
	"context"

	dbmocks "go-backend-template/internal/database/mocks"
	usecasemocks "go-backend-template/internal/usecase/mocks"
	cryptomocks "go-backend-template/internal/util/crypto/mocks"
)

type testPrep struct {
	ctx      context.Context
	crypto   *cryptomocks.Crypto
	config   *usecasemocks.Config
	userRepo *dbmocks.UserRepo

	authUsecases AuthUsecases
	userUsecases UserUsecases
}

func newTestPrep() testPrep {
	crypto := &cryptomocks.Crypto{}
	config := &usecasemocks.Config{}
	userRepo := &dbmocks.UserRepo{}

	db := dbmocks.NewService(userRepo)

	usecases := NewUsecases(db, config, crypto)

	return testPrep{
		ctx:      context.Background(),
		crypto:   crypto,
		config:   config,
		userRepo: userRepo,

		authUsecases: usecases.Auth,
		userUsecases: usecases.User,
	}
}
