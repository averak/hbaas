package global_kvs_usecase

import (
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type Usecase struct {
	conn          transaction.Connection
	globalKVSRepo repository.GlobalKVSRepository
}

func NewUsecase(conn transaction.Connection, globalKVSRepo repository.GlobalKVSRepository) *Usecase {
	return &Usecase{
		conn:          conn,
		globalKVSRepo: globalKVSRepo,
	}
}
