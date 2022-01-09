package repository

import (
	"context"

	"go-backend-template/internal/errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Executor

type QueryExecutor interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

// queries

type Queries struct {
	Executor QueryExecutor
	Builder  goqu.DialectWrapper
}

func NewQueries(executor QueryExecutor) Queries {
	return Queries{
		Executor: executor,
		Builder:  goqu.Dialect("postgres"),
	}
}

func (q *Queries) Rows(ctx context.Context, sql string) (pgx.Rows, error) {
	if q.Executor == nil {
		return nil, errors.New(errors.DatabaseError, "client is not connected")
	}

	rows, err := q.Executor.Query(ctx, sql)
	if err != nil {
		return nil, errors.New(errors.DatabaseError, "").SetInternal(err)
	}

	return rows, nil
}

func (q *Queries) Row(ctx context.Context, sql string) (pgx.Row, error) {
	if q.Executor == nil {
		return nil, errors.New(errors.DatabaseError, "client is not connected")
	}

	return q.Executor.QueryRow(ctx, sql), nil
}

func (q *Queries) Exec(ctx context.Context, sql string) error {
	if q.Executor == nil {
		return errors.New(errors.DatabaseError, "client is not connected")
	}

	_, err := q.Executor.Exec(ctx, sql)
	if err != nil {
		return errors.New(errors.DatabaseError, "").SetInternal(err)
	}

	return nil
}
