package leader_board_usecase

import (
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type Usecase struct {
	conn            transaction.Connection
	leaderBoardRepo repository.LeaderBoardRepository
}

func NewUsecase(conn transaction.Connection, leaderBoardRepo repository.LeaderBoardRepository) *Usecase {
	return &Usecase{
		conn:            conn,
		leaderBoardRepo: leaderBoardRepo,
	}
}
