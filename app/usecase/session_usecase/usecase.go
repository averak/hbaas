package session_usecase

import (
	"context"

	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type Usecase struct {
	conn transaction.Connection

	identityVerifier IdentityVerifier
	authRepo         repository.AuthenticationRepository
	userRepo         repository.UserRepository
}

func NewUsecase(
	conn transaction.Connection,
	identityVerifier IdentityVerifier,
	authRepo repository.AuthenticationRepository,
	userRepo repository.UserRepository,
) *Usecase {
	return &Usecase{
		conn:             conn,
		identityVerifier: identityVerifier,
		authRepo:         authRepo,
		userRepo:         userRepo,
	}
}

type IdentityVerifier interface {
	VerifyIDToken(ctx context.Context, idToken string) (Identity, error)
}

type Identity struct {
	BaasUserID string
	Email      string
}

func NewIdentity(baasUserID, email string) Identity {
	return Identity{
		BaasUserID: baasUserID,
		Email:      email,
	}
}
