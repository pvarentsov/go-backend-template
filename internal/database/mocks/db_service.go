package mocks

import (
	"context"

	"go-backend-template/internal/database"

	mock "github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock

	userRepo *UserRepo
}

func (s *Service) UserRepo() database.UserRepo {
	return s.userRepo
}

func (s *Service) BeginTx(ctx context.Context, do func(ctx context.Context) error) error {
	return do(ctx)
}
