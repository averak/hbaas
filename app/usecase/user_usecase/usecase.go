package user_usecase

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type Usecase struct {
	conn transaction.Connection

	authRepo              repository.AuthenticationRepository
	userRepo              repository.UserRepository
	baasUserDeletionTaskQ BaasUserDeletionTaskQueue
}

func NewUsecase(
	conn transaction.Connection,
	authRepo repository.AuthenticationRepository,
	userRepo repository.UserRepository,
	baasUserDeletionTaskQ BaasUserDeletionTaskQueue,
) *Usecase {
	return &Usecase{
		conn:                  conn,
		userRepo:              userRepo,
		authRepo:              authRepo,
		baasUserDeletionTaskQ: baasUserDeletionTaskQ,
	}
}

type BaasUserDeletionTaskQueue interface {
	Enqueue(ctx context.Context, auth model.UserAuthentication) error
}
