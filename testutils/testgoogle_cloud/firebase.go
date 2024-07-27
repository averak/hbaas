package testgoogle_cloud

import (
	"context"
	"time"

	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/google/uuid"
)

type Option func(*firebaseClientStub)

type firebaseClientStub struct {
	verifyIDTokenFn func(idToken string) (*google_cloud.FirebaseAuthIDToken, error)
	deleteUserFn    func(ctx context.Context, uid string) error
}

func NewFirebaseClientStub(opts ...Option) google_cloud.FirebaseClient {
	cli := firebaseClientStub{
		verifyIDTokenFn: func(idToken string) (*google_cloud.FirebaseAuthIDToken, error) {
			uid := uuid.NewString()
			return &google_cloud.FirebaseAuthIDToken{
				AuthTime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
				IssuedAt: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
				Expires:  time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC).Unix(),
				Subject:  uid,
				UID:      uid,
				Claims: map[string]any{
					google_cloud.FirebaseAuthEmailClaimKey: faker.Email(),
				},
			}, nil
		},
		deleteUserFn: func(ctx context.Context, uid string) error {
			return nil
		},
	}
	for _, opt := range opts {
		opt(&cli)
	}
	return cli
}

func (c firebaseClientStub) VerifyIDToken(ctx context.Context, idToken string) (*google_cloud.FirebaseAuthIDToken, error) {
	return c.verifyIDTokenFn(idToken)
}

func (c firebaseClientStub) DeleteUser(ctx context.Context, uid string) error {
	return c.deleteUserFn(ctx, uid)
}

func WithVerifyIDTokenFn(f func(idToken string) (*google_cloud.FirebaseAuthIDToken, error)) Option {
	return func(c *firebaseClientStub) {
		c.verifyIDTokenFn = f
	}
}

func WithDeleteUserFn(f func(ctx context.Context, uid string) error) Option {
	return func(c *firebaseClientStub) {
		c.deleteUserFn = f
	}
}
