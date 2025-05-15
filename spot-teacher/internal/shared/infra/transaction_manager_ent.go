package infra

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
)

type transactionManagerEnt struct {
	client *ent.Client
}

func NewTransactionManagerEnt(client *ent.Client) TransactionManager {
	return &transactionManagerEnt{client: client}
}

func (m *transactionManagerEnt) Do(
	ctx context.Context,
	fn func(ctx context.Context, r *Repositories) error,
) error {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback() // commit 済みなら harmless
	}()

	repos := &Repositories{}

	if err := fn(ctx, repos); err != nil {
		return err // ロールバックは defer に任せる
	}
	return tx.Commit()
}
