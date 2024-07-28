package private_kvs_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) Get(ctx context.Context, tctx transaction_context.TransactionContext, user model.User, criteria model.KVSCriteria) (model.PrivateKVSBucket, error) {
	var bucket model.PrivateKVSBucket
	err := u.conn.BeginRoTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		var err error
		bucket, err = u.privateKVSRepo.Get(ctx, tx, user.ID, criteria)
		if err != nil {
			return fmt.Errorf("privateKVSRepo.Get failed: %w", err)
		}
		return nil
	})
	if err != nil {
		return model.PrivateKVSBucket{}, err
	}
	return bucket, nil
}
