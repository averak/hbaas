package repository

import (
	"context"
	"errors"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

var (
	ErrMasterDataNotFound = errors.New("master data not found")
)

type MasterDataRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, revision int) (model.MasterData, error)
	GetRevisions(ctx context.Context, tx transaction.Transaction) ([]int, error)
	GetActive(ctx context.Context, tx transaction.Transaction) (model.MasterData, error)
	Save(ctx context.Context, tx transaction.Transaction, masterData model.MasterData) error
}
