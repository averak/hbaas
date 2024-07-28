package leader_board_repoimpl

import (
	"context"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/infrastructure/trace"
	"github.com/averak/hbaas/pkg/vector"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct{}

func NewRepository() repository.LeaderBoardRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, id string) (model.LeaderBoard, error) {
	ctx, span := trace.StartSpan(ctx, "leader_board_repoimpl.Get")
	defer span.End()

	exists, err := dao.LeaderBoards(dao.LeaderBoardWhere.ID.EQ(id)).Exists(ctx, tx)
	if err != nil {
		return model.LeaderBoard{}, err
	}
	if !exists {
		return model.NewLeaderBoard(id, nil), nil
	}

	dtos, err := dao.LeaderBoardScores(dao.LeaderBoardScoreWhere.LeaderBoardID.EQ(id)).All(ctx, tx)
	if err != nil {
		return model.LeaderBoard{}, err
	}
	scores := vector.Map(dtos, func(dto *dao.LeaderBoardScore) model.LeaderBoardScore {
		return model.NewLeaderBoardScore(dto.ScoreID, dto.Score, dto.Timestamp)
	})
	return model.NewLeaderBoard(id, scores), nil
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, leaderBoard model.LeaderBoard) error {
	ctx, span := trace.StartSpan(ctx, "leader_board_repoimpl.Save")
	defer span.End()

	dto := &dao.LeaderBoard{
		ID: leaderBoard.ID,
	}
	err := dto.Upsert(ctx, tx, true, dao.LeaderBoardPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	// 既存のスコアを全削除してから、新しく挿入する。
	// パフォーマンスに懸念があるが、要件的に更新頻度は低いため問題ないと判断した。
	_, err = dao.LeaderBoardScores(dao.LeaderBoardScoreWhere.LeaderBoardID.EQ(leaderBoard.ID)).DeleteAll(ctx, tx)
	if err != nil {
		return err
	}
	scores := vector.Map(leaderBoard.Scores, func(score model.LeaderBoardScore) *dao.LeaderBoardScore {
		return &dao.LeaderBoardScore{
			LeaderBoardID: leaderBoard.ID,
			ScoreID:       score.ScoreID,
			Score:         score.Score,
			Timestamp:     score.Timestamp,
		}
	})
	_, err = dao.LeaderBoardScoreSlice(scores).InsertAll(ctx, tx, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
