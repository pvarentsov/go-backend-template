package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"

	"go-backend-template/internal/util/errors"
)

// Service

type Service interface {
	BeginTx(ctx context.Context, do func(ctx context.Context) error) error
	UserRepo() UserRepo
}

func NewService(client *Client) Service {
	repos := newRepos(client.pool, goqu.Dialect("postgres"))

	return &service{
		client:   client,
		userRepo: repos.User,
	}
}

type service struct {
	client   *Client
	userRepo UserRepo
}

func (s *service) UserRepo() UserRepo {
	return s.userRepo
}

func (s *service) BeginTx(ctx context.Context, do func(ctx context.Context) error) error {
	_, ok := hasTx(ctx)
	if ok {
		return do(ctx)
	}

	return s.beginTx(ctx, do)
}

func (s *service) beginTx(ctx context.Context, do func(ctx context.Context) error) error {
	conn, err := s.client.pool.Begin(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot open transaction").SetInternal(err)
	}

	tx := newTransaction(conn)
	txCtx := withTx(ctx, tx)

	err = do(txCtx)
	if err != nil {
		if err := tx.rollback(txCtx); err != nil {
			return err
		}
		return err
	}
	if err := tx.commit(txCtx); err != nil {
		return err
	}

	return nil
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

func (t *Transaction) commit(ctx context.Context) error {
	err := t.conn.Commit(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot commit transaction").SetInternal(err)
	}

	return nil
}

func (t *Transaction) rollback(ctx context.Context) error {
	err := t.conn.Rollback(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot rollback transaction").SetInternal(err)
	}

	return nil
}

// Transaction Context

type txKey = int

const (
	key txKey = iota
)

func withTx(ctx context.Context, tx Transaction) context.Context {
	return context.WithValue(ctx, key, tx)
}

func hasTx(ctx context.Context) (Transaction, bool) {
	tx, ok := ctx.Value(key).(Transaction)
	if ok {
		return tx, true
	}

	return Transaction{}, false
}
