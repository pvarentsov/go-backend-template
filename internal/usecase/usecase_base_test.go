package usecase

import (
	"context"

	dbmocks "go-backend-template/internal/database/mocks"
	usecasemocks "go-backend-template/internal/usecase/mocks"
	cryptomocks "go-backend-template/internal/util/crypto/mocks"
)

type testUsecase struct {
	ctx      context.Context
	crypto   *cryptomocks.Crypto
	userRepo *dbmocks.UserRepo

	userUsecases UserUsecases
}

func newTestUsecases() testUsecase {
	crypto := &cryptomocks.Crypto{}
	userRepo := &dbmocks.UserRepo{}
	db := dbmocks.NewService(userRepo)

	config := &usecasemocks.Config{}
	usecases := NewUsecases(db, config, crypto)

	return testUsecase{
		ctx:      context.Background(),
		crypto:   crypto,
		userRepo: userRepo,

		userUsecases: usecases.User,
	}
}
