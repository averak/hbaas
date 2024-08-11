package userup

import (
	"context"
	"testing"

	"github.com/averak/hbaas/app/adapter/repoimpl/authentication_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/private_kvs_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_profile_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_repoimpl"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/fixture/builder/user_builder"
)

var repos = struct {
	UserRepo        repository.UserRepository
	UserProfileRepo repository.UserProfileRepository
	AuthRepo        repository.AuthenticationRepository
	PrivateKVSRepo  repository.PrivateKVSRepository
}{
	UserRepo:        user_repoimpl.NewRepository(),
	UserProfileRepo: user_profile_repoimpl.NewRepository(),
	AuthRepo:        authentication_repoimpl.NewRepository(),
	PrivateKVSRepo:  private_kvs_repoimpl.NewRepository(),
}

func Setup(t *testing.T, ctx context.Context, data ...user_builder.Data) {
	t.Helper()

	if len(data) == 0 {
		return
	}

	conn := testutils.MustDBConn(t)
	err := conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		for _, d := range data {
			err := repos.UserRepo.Save(ctx, tx, d.User)
			if err != nil {
				return err
			}

			if d.Profile != nil {
				err = repos.UserProfileRepo.Save(ctx, tx, *d.Profile)
				if err != nil {
					return err
				}
			}

			if d.Authentication != nil {
				err = repos.AuthRepo.Save(ctx, tx, *d.Authentication)
				if err != nil {
					return err
				}
			}

			if d.PrivateKVSBucket != nil {
				err = repos.PrivateKVSRepo.Save(ctx, tx, *d.PrivateKVSBucket)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
