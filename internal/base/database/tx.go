package database

import (
	"context"

	"github.com/jackc/pgx/v4"

	"go-backend-template/internal/base/errors"
)

type txKey = int

const (
	key txKey = iota
)

type transaction struct {
	conn pgx.Tx
}

func (t *transaction) commit(ctx context.Context) error {
	err := t.conn.Commit(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot commit transaction").SetInternal(err)
	}

	return nil
}

func (t *transaction) rollback(ctx context.Context) error {
	err := t.conn.Rollback(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot rollback transaction").SetInternal(err)
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
		return errors.New(errors.DatabaseError, "cannot open transaction").SetInternal(err)
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
