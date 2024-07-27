package google_cloud

import (
	"context"
	"errors"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/infrastructure/trace"
)

type (
	FirebaseAuthIDToken auth.Token
)

var (
	ErrFirebaseAuthInvalidIDToken = errors.New("firebase auth id token is invalid")
	ErrFirebaseAuthExpiredIDToken = errors.New("firebase auth id token is expired")
	ErrFirebaseAuthUserNotFound   = errors.New("firebase auth user not found")
)

const (
	FirebaseAuthEmailClaimKey = "email"
)

type FirebaseClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (*FirebaseAuthIDToken, error)
	DeleteUser(ctx context.Context, uid string) error
}

func NewFirebaseClient(ctx context.Context) (FirebaseClient, error) {
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: config.Get().GetGoogleCloud().GetProjectId()})
	if err != nil {
		return nil, err
	}
	authCli, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return firebaseClient{authCli}, nil
}

type firebaseClient struct {
	auth *auth.Client
}

func (c firebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*FirebaseAuthIDToken, error) {
	ctx, span := trace.StartSpan(ctx, "baas.VerifyIDToken")
	defer span.End()

	result, err := c.auth.VerifyIDToken(ctx, idToken)
	if err != nil {
		if auth.IsIDTokenInvalid(err) {
			return nil, fmt.Errorf("auth.VerifyIDToken failed: %w, %w", err, ErrFirebaseAuthInvalidIDToken)
		}
		if auth.IsIDTokenExpired(err) {
			return nil, fmt.Errorf("auth.VerifyIDToken failed: %w, %w", err, ErrFirebaseAuthExpiredIDToken)
		}
		return nil, fmt.Errorf("auth.VerifyIDToken failed: %w", err)
	}
	return (*FirebaseAuthIDToken)(result), nil
}

func (c firebaseClient) DeleteUser(ctx context.Context, uid string) error {
	ctx, span := trace.StartSpan(ctx, "baas.DeleteUser")
	defer span.End()

	err := c.auth.DeleteUser(ctx, uid)
	if err != nil {
		if auth.IsUserNotFound(err) {
			return fmt.Errorf("auth.DeleteUser failed: %w, %w", err, ErrFirebaseAuthUserNotFound)
		}
		return fmt.Errorf("auth.DeleteUser failed: %w", err)
	}
	return nil
}

func (t FirebaseAuthIDToken) Email() (string, bool) {
	email, ok := t.Claims[FirebaseAuthEmailClaimKey].(string)
	return email, ok
}
