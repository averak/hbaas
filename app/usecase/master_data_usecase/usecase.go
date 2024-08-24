package master_data_usecase

import (
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type Usecase struct {
	conn           transaction.Connection
	masterDataRepo repository.MasterDataRepository
}

func NewUsecase(conn transaction.Connection, masterDataRepo repository.MasterDataRepository) *Usecase {
	return &Usecase{
		conn:           conn,
		masterDataRepo: masterDataRepo,
	}
}
