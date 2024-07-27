package usecaseimpl

import (
	"context"
	"fmt"

	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"github.com/averak/hbaas/app/usecase/session_usecase"
)

type FirebaseIdentityVerifier struct {
	cli google_cloud.FirebaseClient
}

func NewFirebaseIdentityVerifier(cli google_cloud.FirebaseClient) session_usecase.IdentityVerifier {
	return FirebaseIdentityVerifier{cli}
}

func (f FirebaseIdentityVerifier) VerifyIDToken(ctx context.Context, idToken string) (session_usecase.Identity, error) {
	token, err := f.cli.VerifyIDToken(ctx, idToken)
	if err != nil {
		return session_usecase.Identity{}, fmt.Errorf("VerifyIDToken failed: %w", err)
	}
	email, ok := token.Email()
	if !ok {
		return session_usecase.Identity{}, fmt.Errorf("email not found in id token")
	}
	return session_usecase.NewIdentity(token.UID, email), nil
}
