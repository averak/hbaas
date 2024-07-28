package global_kvs_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) Get(ctx context.Context, criteria model.KVSCriteria) (model.GlobalKVSBucket, error) {
	var bucket model.GlobalKVSBucket
	err := u.conn.BeginRoTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		var err error
		bucket, err = u.globalKVSRepo.Get(ctx, tx, criteria)
		if err != nil {
			return fmt.Errorf("globalKVSRepo.Get failed: %w", err)
		}
		return nil
	})
	if err != nil {
		return model.GlobalKVSBucket{}, err
	}
	return bucket, nil
}
