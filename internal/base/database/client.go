package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"go-backend-template/internal/base/errors"
)

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
