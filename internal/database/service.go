package database

import (
	"context"

	"go-backend-template/internal/database/repository"
	"go-backend-template/internal/errors"
)

type Service interface {
	Repositories() repository.Repositories
	BeginTx(ctx context.Context) (Transaction, error)
}

func NewService(client *Client) Service {
	queries := repository.NewQueries(client.pool)
	repositories := repository.NewRepositories(queries)

	service := service{
		repositories: repositories,
		client:       client,
	}

	return &service
}

type service struct {
	repositories repository.Repositories
	client       *Client
}

func (s *service) Repositories() repository.Repositories {
	return s.repositories
}

func (s *service) BeginTx(ctx context.Context) (Transaction, error) {
	if s.client.pool == nil {
		return nil, errors.New(errors.DatabaseError, "client is not connected")
	}

	executor, err := s.client.pool.Begin(ctx)
	if err != nil {
		return nil, errors.New(errors.DatabaseError, "cannot open transaction").SetInternal(err)
	}

	queries := repository.NewQueries(executor)
	repositories := repository.NewRepositories(queries)
	transaction := NewTransaction(executor, repositories)

	return transaction, nil
}
