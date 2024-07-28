package repository

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type LeaderBoardRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, id string) (model.LeaderBoard, error)
	Save(ctx context.Context, tx transaction.Transaction, leaderBoard model.LeaderBoard) error
}
