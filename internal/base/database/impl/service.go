package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	"go-backend-template/internal/base/errors"
)

type ConnManager interface {
	Conn(ctx context.Context) Connection
}

type Connection interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

func NewService(client *Client) *Service {
	return &Service{
		client: client,
	}
}

type Service struct {
	client *Client
}

func (s *Service) RunTx(ctx context.Context, do func(ctx context.Context) error) error {
	_, ok := hasTx(ctx)
	if ok {
		return do(ctx)
	}

	return runTx(ctx, s.client, do)
}

func (s *Service) Conn(ctx context.Context) Connection {
	tx, ok := hasTx(ctx)
	if ok {
		return tx.conn
	}

	return s.client.pool
}

type txKey int

const (
	key txKey = iota
)

type transaction struct {
	conn pgx.Tx
}

func (t *transaction) commit(ctx context.Context) error {
	err := t.conn.Commit(ctx)
	if err != nil {
		return errors.Wrap(errors.DatabaseError, err, "cannot commit transaction")
	}

	return nil
}

func (t *transaction) rollback(ctx context.Context) error {
	err := t.conn.Rollback(ctx)
	if err != nil {
		return errors.Wrap(errors.DatabaseError, err, "cannot rollback transaction")
	}

	return nil
}

func withTx(ctx context.Context, tx transaction) context.Context {
	return context.WithValue(ctx, key, tx)
}

func hasTx(ctx context.Context) (transaction, bool) {
	tx, ok := ctx.Value(key).(transaction)
	if ok {
		return tx, true
	}

	return transaction{}, false
}

func runTx(ctx context.Context, client *Client, do func(ctx context.Context) error) error {
	conn, err := client.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(errors.DatabaseError, err, "cannot open transaction")
	}

	tx := transaction{conn: conn}
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

var QueryBuilder = goqu.Dialect("postgres")

type Ex = goqu.Ex
type Record = goqu.Record
