package session_usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/google/uuid"
)

type AuthorizeResult struct {
	UserID       uuid.UUID
	AuthorizedAt time.Time
}

// Authorize はユーザを認証し、セッションを新規作成します。
// 初めてセッションを作成するユーザの場合は、アカウントの新規作成も行います。
func (u Usecase) Authorize(ctx context.Context, tctx transaction_context.TransactionContext, idToken string) (AuthorizeResult, error) {
	identity, err := u.identityVerifier.VerifyIDToken(ctx, idToken)
	if err != nil {
		return AuthorizeResult{}, fmt.Errorf("identityVerifier.VerifyIDToken failed, %w", err)
	}

	var auth model.UserAuthentication
	err = u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		auth, err = u.authRepo.GetByBaasUserID(ctx, tx, identity.BaasUserID)
		if errors.Is(err, repository.ErrUserAuthenticationNotFound) {
			// 認証情報が存在しない => 存在しないユーザとして、サインアップする。
			user := model.NewUser(uuid.New(), identity.Email, model.UserStatusPending)
			err = u.userRepo.Save(ctx, tx, user)
			if err != nil {
				return fmt.Errorf("userRepo.Save failed, %w", err)
			}
			auth = model.NewUserAuthentication(user.ID, identity.BaasUserID, tctx.Now())
		} else if err != nil {
			return fmt.Errorf("authRepo.GetByBaasUserID failed, %w", err)
		} else {
			// サインアップ時にリトライされたらここを通るので、冪等に実行できる。
			// したがって、想定外のリトライによって複数ユーザ作成されることはない。
			auth.Refresh(tctx.Now())
		}

		err = u.authRepo.Save(ctx, tx, auth)
		if err != nil {
			return fmt.Errorf("authRepo.Save failed, %w", err)
		}
		return nil
	})
	if err != nil {
		return AuthorizeResult{}, err
	}
	return AuthorizeResult{auth.UserID, auth.LastAuthenticatedAt}, nil
}
