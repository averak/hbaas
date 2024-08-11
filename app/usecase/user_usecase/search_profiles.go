package user_usecase

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/pkg/vector"
	"github.com/google/uuid"
)

func (u Usecase) SearchProfiles(ctx context.Context, userIDs []uuid.UUID) ([]model.UserProfile, error) {
	var profiles []model.UserProfile
	err := u.conn.BeginRoTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		var err error
		users, err := u.userRepo.GetByUserIDs(ctx, tx, userIDs)
		if err != nil {
			return fmt.Errorf("userRepo.GetByUserIDs failed, %w", err)
		}

		activeUserIDs := vector.Map(
			vector.New(users).Filter(func(user model.User) bool {
				return user.Status == model.UserStatusActive
			}),
			func(user model.User) uuid.UUID {
				return user.ID
			},
		)
		profiles, err = u.userProfileRepo.GetByUserIDs(ctx, tx, activeUserIDs)
		if err != nil {
			return fmt.Errorf("userProfileRepo.GetByUserIDs failed, %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return profiles, nil
}
