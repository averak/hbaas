package global_kvs_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/pkg/vector"
)

func (u Usecase) Set(ctx context.Context, entries []model.KVSEntry) error {
	criteria := model.NewKVSCriteria(
		vector.Map(entries, func(entry model.KVSEntry) string {
			return entry.Key
		}),
		nil,
	)
	err := u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		bucket, err := u.globalKVSRepo.Get(ctx, tx, criteria)
		if err != nil {
			return fmt.Errorf("globalKVSRepo.Get failed: %w", err)
		}

		bucket.Set(entries)
		err = u.globalKVSRepo.Save(ctx, tx, bucket)
		if err != nil {
			return fmt.Errorf("globalKVSRepo.Save failed: %w", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
