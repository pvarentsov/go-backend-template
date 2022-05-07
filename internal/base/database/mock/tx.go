package mock

import "context"

type MockTxManager struct{}

func (*MockTxManager) RunTx(ctx context.Context, do func(ctx context.Context) error) error {
	return do(ctx)
}
