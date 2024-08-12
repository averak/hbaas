package master_data_usecase

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) Get(ctx context.Context) (model.MasterData, error) {
	var res model.MasterData
	err := u.conn.BeginRoTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		var err error
		res, err = u.masterDataRepo.GetActive(ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return model.MasterData{}, err
	}
	return res, nil
}
