package user_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) EditProfile(ctx context.Context, user model.User, data []byte) error {
	profile, err := model.NewUserProfile(user.ID, data)
	if err != nil {
		return err
	}
	err = u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		err = u.userProfileRepo.Save(ctx, tx, profile)
		if err != nil {
			return fmt.Errorf("userProfileRepo.Save failed, %w", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
