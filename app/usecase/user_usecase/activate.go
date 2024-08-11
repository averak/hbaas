package user_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) Activate(ctx context.Context, user model.User) error {
	err := user.Activate()
	if err != nil {
		return err
	}
	err = u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		err = u.userRepo.Save(ctx, tx, user)
		if err != nil {
			return fmt.Errorf("userRepo.Save failed, %w", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
