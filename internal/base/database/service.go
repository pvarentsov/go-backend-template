package database

import (
	"context"
)

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
