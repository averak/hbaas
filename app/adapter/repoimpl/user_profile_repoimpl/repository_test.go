package user_profile_repoimpl

import (
	"context"
	"testing"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/averak/hbaas/testutils/fixture"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_Get(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type args struct {
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.UserProfile
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "レコードが存在する => 取得できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
				&dao.UserProfile{
					UserID:  faker.UUIDv5("u1").String(),
					Content: []byte("value"),
				},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
			},
			want:    mustUserProfile(t, faker.UUIDv5("u1"), []byte("value")),
			wantErr: assert.NoError,
		},
		{
			name:  "レコードが存在しない => デフォルト値を返す",
			seeds: []fixture.Seed{},
			args: args{
				userID: faker.UUIDv5("u1"),
			},
			want:    mustUserProfile(t, faker.UUIDv5("u1"), nil),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			var got model.UserProfile
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
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRepository_GetByUserIDs(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type args struct {
		userIDs []uuid.UUID
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    []model.UserProfile
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ユーザIDで検索できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID:    faker.UUIDv5("u1").String(),
					Email: faker.Email(),
				},
				&dao.UserProfile{
					UserID:  faker.UUIDv5("u1").String(),
					Content: []byte("v1"),
				},
				&dao.User{
					ID:    faker.UUIDv5("u2").String(),
					Email: faker.Email(),
				},
				&dao.UserProfile{
					UserID:  faker.UUIDv5("u2").String(),
					Content: []byte("v2"),
				},
				&dao.User{
					ID:    faker.UUIDv5("u3").String(),
					Email: faker.Email(),
				},
				&dao.UserProfile{
					UserID:  faker.UUIDv5("u3").String(),
					Content: []byte("v3"),
				},
			},
			args: args{
				userIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
					faker.UUIDv5("u2"),

					// 存在しないユーザは無視される
					faker.UUIDv5("not_exists"),
				},
			},
			want: []model.UserProfile{
				mustUserProfile(t, faker.UUIDv5("u1"), []byte("v1")),
				mustUserProfile(t, faker.UUIDv5("u2"), []byte("v2")),
			},
			wantErr: assert.NoError,
		},
		{
			name: "ユーザIDリストが空 => 空リストを返す",
			args: args{
				userIDs: []uuid.UUID{},
			},
			want:    []model.UserProfile{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			var got []model.UserProfile
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := Repository{}
				var err error
				got, err = r.GetByUserIDs(ctx, tx, tt.args.userIDs)
				if err != nil {
					return err
				}
				return nil
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRepository_Save(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type args struct {
		profile model.UserProfile
	}
	tests := []struct {
		name  string
		seeds []fixture.Seed
		args  args
		then  func(*testing.T, []*dao.UserProfile, error)
	}{
		{
			name: "PK が存在しない => 作成する",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
			},
			args: args{
				profile: mustUserProfile(t, faker.UUIDv5("u1"), []byte("value")),
			},
			then: func(t *testing.T, dtos []*dao.UserProfile, err error) {
				require.NoError(t, err)

				wantDtos := []*dao.UserProfile{
					{
						UserID:  faker.UUIDv5("u1").String(),
						Content: []byte("value"),
					},
				}
				if diff := cmp.Diff(wantDtos, dtos, cmpopts.IgnoreFields(dao.UserProfile{}, "CreatedAt", "UpdatedAt")); diff != "" {
					t.Errorf("(-want, +got)\n%s", diff)
				}
			},
		},
		{
			name: "PK が存在する => 更新する",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
				&dao.UserProfile{
					UserID:  faker.UUIDv5("u1").String(),
					Content: []byte("old"),
				},
			},
			args: args{
				profile: mustUserProfile(t, faker.UUIDv5("u1"), []byte("new")),
			},
			then: func(t *testing.T, dtos []*dao.UserProfile, err error) {
				require.NoError(t, err)

				wantDtos := []*dao.UserProfile{
					{
						UserID:  faker.UUIDv5("u1").String(),
						Content: []byte("new"),
					},
				}
				if diff := cmp.Diff(wantDtos, dtos, cmpopts.IgnoreFields(dao.UserProfile{}, "CreatedAt", "UpdatedAt")); diff != "" {
					t.Errorf("(-want, +got)\n%s", diff)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			var dtos []*dao.UserProfile
			err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := Repository{}
				err := r.Save(ctx, tx, tt.args.profile)
				if err != nil {
					return err
				}

				dtos, err = dao.UserProfiles().All(ctx, tx)
				if err != nil {
					return err
				}
				return nil
			})
			tt.then(t, dtos, err)
		})
	}
}

func mustUserProfile(t *testing.T, userID uuid.UUID, content []byte) model.UserProfile {
	profile, err := model.NewUserProfile(userID, content)
	if err != nil {
		t.Fatal(err)
	}
	return profile
}
