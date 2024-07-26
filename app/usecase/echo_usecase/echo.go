package echo_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) Echo(ctx context.Context, tctx transaction_context.TransactionContext, message string) (model.Echo, error) {
	echo := model.NewEcho(message, tctx.Now())
	err := u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		err := u.echoRepo.Save(ctx, tx, echo)
		if err != nil {
			return fmt.Errorf("echoRepo.Save failed, %w", err)
		}
		return nil
	})
	if err != nil {
		return model.Echo{}, err
	}
	return echo, nil
}
