package authentication_repoimpl

import (
	"context"
	"testing"
	"time"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/averak/hbaas/testutils/fixture"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Get(t *testing.T) {
	conn := testutils.MustDBConn(t)
	now := time.Now().Truncate(time.Millisecond)

	type args struct {
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.UserAuthentication
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "レコードが存在する => 取得できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
				&dao.UserAuthentication{
					UserID:              faker.UUIDv5("u1").String(),
					BaasUserID:          faker.UUIDv5("bu1").String(),
					LastAuthenticatedAt: now,
				},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
			},
			want: model.UserAuthentication{
				UserID:              faker.UUIDv5("u1"),
				BaasUserID:          faker.UUIDv5("bu1").String(),
				LastAuthenticatedAt: now,
			},
			wantErr: assert.NoError,
		},
		{
			name: "レコードが存在しない => エラー",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, repository.ErrUserAuthenticationNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			var got model.UserAuthentication
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := Repository{}
				var err error
				got, err = r.Get(ctx, tx, tt.args.userID)
				if err != nil {
					return err
				}
				return nil
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.EqualExportedValues(t, tt.want, got)
		})
	}
}

func TestRepository_GetByBaasUserID(t *testing.T) {
	conn := testutils.MustDBConn(t)
	now := time.Now().Truncate(time.Millisecond)

	type args struct {
		baasUserID string
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.UserAuthentication
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "レコードが存在する => 取得できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
				&dao.UserAuthentication{
					UserID:              faker.UUIDv5("u1").String(),
					BaasUserID:          faker.UUIDv5("bu1").String(),
					LastAuthenticatedAt: now,
				},
			},
			args: args{
				baasUserID: faker.UUIDv5("bu1").String(),
			},
			want: model.UserAuthentication{
				UserID:              faker.UUIDv5("u1"),
				BaasUserID:          faker.UUIDv5("bu1").String(),
				LastAuthenticatedAt: now,
			},
			wantErr: assert.NoError,
		},
		{
			name: "レコードが存在しない => エラー",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
			},
			args: args{
				baasUserID: faker.UUIDv5("bu1").String(),
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, repository.ErrUserAuthenticationNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			var got model.UserAuthentication
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := Repository{}
				var err error
				got, err = r.GetByBaasUserID(ctx, tx, tt.args.baasUserID)
				if err != nil {
					return err
				}
				return nil
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.EqualExportedValues(t, tt.want, got)
		})
	}
}

func TestRepository_Save(t *testing.T) {
	conn := testutils.MustDBConn(t)
	now := time.Now().Truncate(time.Millisecond)

	type args struct {
		auth model.UserAuthentication
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.UserAuthentication
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "PK が存在しない => 作成する",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
			},
			args: args{
				auth: model.UserAuthentication{
					UserID:              faker.UUIDv5("u1"),
					BaasUserID:          faker.UUIDv5("bu1").String(),
					LastAuthenticatedAt: now,
				},
			},
			want: model.UserAuthentication{
				UserID:              faker.UUIDv5("u1"),
				BaasUserID:          faker.UUIDv5("bu1").String(),
				LastAuthenticatedAt: now,
			},
			wantErr: assert.NoError,
		},
		{
			name: "PK が存在する => 更新する",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
				&dao.UserAuthentication{
					UserID: faker.UUIDv5("u1").String(),
				},
			},
			args: args{
				auth: model.UserAuthentication{
					UserID:              faker.UUIDv5("u1"),
					BaasUserID:          faker.UUIDv5("bu1").String(),
					LastAuthenticatedAt: now,
				},
			},
			want: model.UserAuthentication{
				UserID:              faker.UUIDv5("u1"),
				BaasUserID:          faker.UUIDv5("bu1").String(),
				LastAuthenticatedAt: now,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			var got model.UserAuthentication
			err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := Repository{}
				err := r.Save(ctx, tx, tt.args.auth)
				if err != nil {
					return err
				}

				got, err = r.Get(ctx, tx, tt.args.auth.UserID)
				if err != nil {
					return err
				}
				return nil
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.EqualExportedValues(t, tt.args.auth, got)
		})
	}
}
