package repository

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type GlobalKVSRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, criteria model.KVSCriteria) (model.GlobalKVSBucket, error)
	Save(ctx context.Context, tx transaction.Transaction, bucket model.GlobalKVSBucket) error
}
