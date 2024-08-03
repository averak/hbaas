package repository

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/google/uuid"
)

type UserProfileRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, userID uuid.UUID) (model.UserProfile, error)
	Save(ctx context.Context, tx transaction.Transaction, profile model.UserProfile) error
}
