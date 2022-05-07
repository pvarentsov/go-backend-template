package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var QueryBuilder = goqu.Dialect("postgres")

type Config interface {
	ConnString() string
}

type TxManager interface {
	RunTx(ctx context.Context, do func(ctx context.Context) error) error
}

type ConnManager interface {
	Conn(ctx context.Context) Connection
}

type Connection interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}
