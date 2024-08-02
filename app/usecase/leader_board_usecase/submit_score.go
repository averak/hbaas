package leader_board_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) SubmitScore(ctx context.Context, tctx transaction_context.TransactionContext, leaderBoardID string, scoreID string, score int) (model.LeaderBoard, error) {
	var leaderBoard model.LeaderBoard
	err := u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		var err error
		leaderBoard, err = u.leaderBoardRepo.Get(ctx, tx, leaderBoardID)
		if err != nil {
			return fmt.Errorf("leaderBoardRepo.Get failed: %w", err)
		}

		leaderBoard.SubmitScore(model.NewLeaderBoardScore(scoreID, score, tctx.Now()))
		err = u.leaderBoardRepo.Save(ctx, tx, leaderBoard)
		if err != nil {
			return fmt.Errorf("leaderBoardRepo.Save failed: %w", err)
		}
		return nil
	})
	if err != nil {
		return model.LeaderBoard{}, err
	}
	return leaderBoard, nil
}
