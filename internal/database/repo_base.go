package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Repo

type repo struct {
	pool connection
	qb   goqu.DialectWrapper
}

func (r *userRepo) conn(ctx context.Context) connection {
	tx, ok := hasTx(ctx)
	if ok {
		return tx.conn
	}

	return r.pool
}

// Repos

type Repos struct {
	User UserRepo
}

func newRepos(conn connection, qb goqu.DialectWrapper) Repos {
	return Repos{User: newUserRepo(conn, qb)}
}

// Connection

type connection interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}
