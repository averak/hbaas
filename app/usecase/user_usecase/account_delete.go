package user_usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/averak/hbaas/app/core/retry"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

func (u Usecase) AccountDelete(ctx context.Context, user model.User) error {
	err := u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		user.Delete()
		err := u.userRepo.Save(ctx, tx, user)
		if err != nil {
			return fmt.Errorf("userRepo.Save failed, %w", err)
		}

		// 認証情報は削除しないので、リトライ時も冪等に実行できる。
		auth, err := u.authRepo.Get(ctx, tx, user.ID)
		if err != nil {
			return fmt.Errorf("authRepo.Get failed, %w", err)
		}

		// TODO: ユーザに紐づくのリソースを削除する。

		err = retry.Do(ctx, func() error {
			err = u.baasUserDeletionTaskQ.Enqueue(ctx, auth)
			if err != nil {
				return fmt.Errorf("baasUserDeletionTaskQ.Enqueue failed, %w", err)
			}
			return nil
		}, retry.WithExponentialBackoff(100*time.Millisecond, 800*time.Millisecond), retry.WithJitter(1*time.Second))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
