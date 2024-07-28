package private_kvs_usecase

import (
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type Usecase struct {
	conn           transaction.Connection
	privateKVSRepo repository.PrivateKVSRepository
}

func NewUsecase(conn transaction.Connection, privateKVSRepo repository.PrivateKVSRepository) *Usecase {
	return &Usecase{
		conn:           conn,
		privateKVSRepo: privateKVSRepo,
	}
}
