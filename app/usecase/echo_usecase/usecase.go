package echo_usecase

import (
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type Usecase struct {
	conn     transaction.Connection
	echoRepo repository.EchoRepository
}

func NewUsecase(conn transaction.Connection, echoRepo repository.EchoRepository) *Usecase {
	return &Usecase{
		conn:     conn,
		echoRepo: echoRepo,
	}
}
