package repository

import (
	"context"
	"errors"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/google/uuid"
)

var (
	ErrUserAuthenticationNotFound = errors.New("user authentication not found")
)

type AuthenticationRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, userID uuid.UUID) (model.UserAuthentication, error)
	GetByBaasUserID(ctx context.Context, tx transaction.Transaction, baasUserID string) (model.UserAuthentication, error)
	Save(ctx context.Context, tx transaction.Transaction, auth model.UserAuthentication) error
}
