package repository

import (
	"context"
	"errors"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

var (
	ErrLeaderBoardEventNotFound = errors.New("leader board event not found")
)

type LeaderBoardRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, eventID string) (model.LeaderBoard, error)
	Save(ctx context.Context, tx transaction.Transaction, leaderBoard model.LeaderBoard) error
}
