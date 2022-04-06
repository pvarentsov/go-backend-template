package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"go-backend-template/internal/util/contexts"
	"go-backend-template/internal/util/errors"
)

// Config

type Config interface {
	ConnString() string
}

// Client

type Client struct {
	pool *pgxpool.Pool
	url  string
	ctx  context.Context
}

func NewClient(ctx context.Context, config Config) *Client {
	return &Client{
		ctx: ctx,
		url: config.ConnString(),
	}
}

func (c *Client) Connect() error {
	c.Close()

	config, err := pgxpool.ParseConfig(c.url)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot connect to database").SetInternal(err)
	}

	config.ConnConfig.Logger = &logger{}

	pool, err := pgxpool.ConnectConfig(c.ctx, config)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot connect to database").SetInternal(err)
	}

	c.pool = pool

	return nil
}

func (c *Client) Close() {
	if c.pool != nil {
		c.pool.Close()
	}
}

// Logger

type logger struct{}

func (l *logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	reqInfo, _ := contexts.GetReqInfo(ctx)

	if sql, ok := data["sql"]; ok {
		fmt.Printf("%s - [Database] TraceId: %s; UserId: %d; SQL: %s;\n\n",
			time.Now().Format(time.RFC1123),
			reqInfo.TraceId,
			reqInfo.UserId,
			sql,
		)
	}
}
