package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"

	"go-backend-template/internal/util/errors"
)

// Service

type Service struct {
	UserRepo UserRepo
	client   *Client
}

func NewService(client *Client) Service {
	repos := newRepos(client.pool, goqu.Dialect("postgres"))

	return Service{
		client:   client,
		UserRepo: repos.User,
	}
}

func (s *Service) BeginTx(ctx context.Context) (Transaction, error) {
	conn, err := s.client.pool.Begin(ctx)
	if err != nil {
		return Transaction{}, errors.New(errors.DatabaseError, "cannot open transaction").SetInternal(err)
	}

	transaction := newTransaction(conn)

	return transaction, nil
}

// Transaction

type Transaction struct {
	conn     pgx.Tx
	UserRepo UserRepo
}

func newTransaction(conn pgx.Tx) Transaction {
	repos := newRepos(conn, goqu.Dialect("postgres"))

	return Transaction{
		conn:     conn,
		UserRepo: repos.User,
	}
}

func (t *Transaction) Commit(ctx context.Context) error {
	err := t.conn.Commit(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot commit transaction").SetInternal(err)
	}

	return nil
}

func (t *Transaction) Rollback(ctx context.Context) error {
	err := t.conn.Rollback(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot rollback transaction").SetInternal(err)
	}

	return nil
}
