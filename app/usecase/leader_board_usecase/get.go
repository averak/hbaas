package leader_board_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) Get(ctx context.Context, leaderBoardID string) (model.LeaderBoard, error) {
	var leaderBoard model.LeaderBoard
	err := u.conn.BeginRoTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		var err error
		leaderBoard, err = u.leaderBoardRepo.Get(ctx, tx, leaderBoardID)
		if err != nil {
			return fmt.Errorf("leaderBoardRepo.Get failed: %w", err)
		}
		return nil
	})
	if err != nil {
		return model.LeaderBoard{}, err
	}
	return leaderBoard, nil
}
