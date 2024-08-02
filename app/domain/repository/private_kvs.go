package repository

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/google/uuid"
)

type PrivateKVSRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, userID uuid.UUID, criteria model.KVSCriteria) (model.PrivateKVSBucket, error)
	Save(ctx context.Context, tx transaction.Transaction, bucket model.PrivateKVSBucket) error
}
