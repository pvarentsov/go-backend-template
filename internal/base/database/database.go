package database

import (
	"context"
)

type Config interface {
	ConnString() string
}

type TxManager interface {
	RunTx(ctx context.Context, do func(ctx context.Context) error) error
}
