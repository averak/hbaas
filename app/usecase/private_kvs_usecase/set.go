package private_kvs_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/pkg/vector"
	"github.com/google/uuid"
)

func (u Usecase) Set(ctx context.Context, user model.User, etag uuid.UUID, entries []model.KVSEntry) (model.PrivateKVSBucket, error) {
	criteria := model.NewKVSCriteria(
		vector.Map(entries, func(entry model.KVSEntry) string {
			return entry.Key
		}),
		nil,
	)
	var bucket model.PrivateKVSBucket
	err := u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		var err error
		bucket, err = u.privateKVSRepo.Get(ctx, tx, user.ID, criteria)
		if err != nil {
			return fmt.Errorf("privateKVSRepo.Get failed: %w", err)
		}

		err = bucket.Set(etag, entries)
		if err != nil {
			return err
		}
		err = u.privateKVSRepo.Save(ctx, tx, bucket)
		if err != nil {
			return fmt.Errorf("privateKVSRepo.Save failed: %w", err)
		}
		return nil
	})
	if err != nil {
		return model.PrivateKVSBucket{}, err
	}
	return bucket, nil
}
