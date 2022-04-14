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
	userRepo *dbmocks.UserRepo

	userUsecases UserUsecases
}

func newTestPrep() testPrep {
	crypto := &cryptomocks.Crypto{}
	userRepo := &dbmocks.UserRepo{}
	db := dbmocks.NewService(userRepo)

	config := &usecasemocks.Config{}
	usecases := NewUsecases(db, config, crypto)

	return testPrep{
		ctx:      context.Background(),
		crypto:   crypto,
		userRepo: userRepo,

		userUsecases: usecases.User,
	}
}
