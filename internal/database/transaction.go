package database

import (
	"context"

	"go-backend-template/internal/database/repository"
	"go-backend-template/internal/errors"

	"github.com/jackc/pgx/v4"
)

type Transaction interface {
	Repositories() repository.Repositories
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

func NewTransaction(executor pgx.Tx, repositories repository.Repositories) Transaction {
	return &transaction{
		repositories: repositories,
		executor:     executor,
	}
}

type transaction struct {
	repositories repository.Repositories
	executor     pgx.Tx
}

func (t *transaction) Repositories() repository.Repositories {
	return t.repositories
}

func (t *transaction) Commit(ctx context.Context) error {
	err := t.executor.Commit(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot commit transaction").SetInternal(err)
	}

	return nil
}

func (t *transaction) Rollback(ctx context.Context) error {
	err := t.executor.Rollback(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot rollback transaction").SetInternal(err)
	}

	return nil
}
