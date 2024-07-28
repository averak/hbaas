package private_kvs_repoimpl

import (
	"context"
	"testing"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
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

	type given struct {
		seeds []fixture.Seed
	}
	type when struct {
		userID   uuid.UUID
		criteria model.KVSCriteria
	}
	type then = func(t *testing.T, got model.PrivateKVSBucket, err error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "KVS が初期化済みの状態で",
			Given: given{
				seeds: []fixture.Seed{
					&dao.User{
						ID: faker.UUIDv5("u1").String(),
					},
					&dao.PrivateKVSEntry{
						UserID: faker.UUIDv5("u1").String(),
						Key:    "group1:key1",
						Value:  []byte("v1"),
					},
					&dao.PrivateKVSEntry{
						UserID: faker.UUIDv5("u1").String(),
						Key:    "group1:key2",
						Value:  []byte("v2"),
					},
					&dao.PrivateKVSEntry{
						UserID: faker.UUIDv5("u1").String(),
						Key:    "group2:key1",
						Value:  []byte("v3"),
					},
					&dao.PrivateKVSEntry{
						UserID: faker.UUIDv5("u1").String(),
						Key:    "group2:key2",
						Value:  []byte("v4"),
					},
					&dao.PrivateKVSEntry{
						UserID: faker.UUIDv5("u1").String(),
						Key:    "group3:key1",
						Value:  []byte("v5"),
					},
					&dao.PrivateKVSEtag{
						UserID: faker.UUIDv5("u1").String(),
						Etag:   faker.UUIDv5("e1").String(),
					},
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "前方一致で取得できる",
					When: when{
						userID: faker.UUIDv5("u1"),
						criteria: model.KVSCriteria{
							PrefixMatch: []string{"group1"},
						},
					},
					Then: func(t *testing.T, got model.PrivateKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewPrivateKVSBucket(
							faker.UUIDv5("u1"),
							faker.UUIDv5("e1"),
							[]model.KVSEntry{
								{
									Key:   "group1:key1",
									Value: []byte("v1"),
								},
								{
									Key:   "group1:key2",
									Value: []byte("v2"),
								},
							},
						)
						assert.Equal(t, want, got)
					},
				},
				{
					Name: "完全一致で取得できる",
					When: when{
						userID: faker.UUIDv5("u1"),
						criteria: model.KVSCriteria{
							ExactMatch: []string{
								"group1:key1",
								"group2:key1",
							},
						},
					},
					Then: func(t *testing.T, got model.PrivateKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewPrivateKVSBucket(
							faker.UUIDv5("u1"),
							faker.UUIDv5("e1"),
							[]model.KVSEntry{
								{
									Key:   "group1:key1",
									Value: []byte("v1"),
								},
								{
									Key:   "group2:key1",
									Value: []byte("v3"),
								},
							},
						)
						assert.Equal(t, want, got)
					},
				},
				{
					Name: "前方一致、完全一致の OR 条件で取得できる",
					When: when{
						userID: faker.UUIDv5("u1"),
						criteria: model.KVSCriteria{
							ExactMatch:  []string{"group1:key1"},
							PrefixMatch: []string{"group2"},
						},
					},
					Then: func(t *testing.T, got model.PrivateKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewPrivateKVSBucket(
							faker.UUIDv5("u1"),
							faker.UUIDv5("e1"),
							[]model.KVSEntry{
								{
									Key:   "group1:key1",
									Value: []byte("v1"),
								},
								{
									Key:   "group2:key1",
									Value: []byte("v3"),
								},
								{
									Key:   "group2:key2",
									Value: []byte("v4"),
								},
							},
						)
						assert.Equal(t, want, got)
					},
				},
				{
					Name: "検索条件が空 => 空リストを返す",
					When: when{
						userID:   faker.UUIDv5("u1"),
						criteria: model.KVSCriteria{},
					},
					Then: func(t *testing.T, got model.PrivateKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewPrivateKVSBucket(faker.UUIDv5("u1"), faker.UUIDv5("e1"), nil)
						assert.Equal(t, want, got)
					},
				},
			},
		},
		{
			Name: "KVS が初期化されていない状態で",
			Given: given{
				seeds: []fixture.Seed{},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "ETag が空になる",
					When: when{
						userID:   faker.UUIDv5("u1"),
						criteria: model.KVSCriteria{},
					},
					Then: func(t *testing.T, got model.PrivateKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewPrivateKVSBucket(faker.UUIDv5("u1"), uuid.Nil, nil)
						assert.Equal(t, want, got)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), given.seeds...)

			var got model.PrivateKVSBucket
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				var err error
				got, err = r.Get(ctx, tx, when.userID, when.criteria)
				return err
			})
			then(t, got, err)
		})
	}
}

func TestRepository_Save(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type args struct {
		bucket model.PrivateKVSBucket
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    []*dao.PrivateKVSEntry
		want1   []*dao.PrivateKVSEtag
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "エントリを保存できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
				&dao.PrivateKVSEntry{
					UserID: faker.UUIDv5("u1").String(),
					Key:    "k1",
					Value:  []byte("v1"),
				},
				&dao.PrivateKVSEntry{
					UserID: faker.UUIDv5("u1").String(),
					Key:    "k2",
					Value:  []byte("v2"),
				},
				&dao.PrivateKVSEntry{
					UserID: faker.UUIDv5("u1").String(),
					Key:    "k3",
					Value:  []byte("v3"),
				},
			},
			args: args{
				bucket: model.NewPrivateKVSBucket(
					faker.UUIDv5("u1"),
					faker.UUIDv5("e1"),
					[]model.KVSEntry{
						{ // 更新される
							Key:   "k2",
							Value: []byte("updated v2"),
						},
						{ // 削除される
							Key:   "k3",
							Value: nil,
						},
						{ // 作成される
							Key:   "k4",
							Value: []byte("inserted v4"),
						},
					},
				),
			},
			want: []*dao.PrivateKVSEntry{
				{
					UserID: faker.UUIDv5("u1").String(),
					Key:    "k1",
					Value:  []byte("v1"),
				},
				{
					UserID: faker.UUIDv5("u1").String(),
					Key:    "k2",
					Value:  []byte("updated v2"),
				},
				{
					UserID: faker.UUIDv5("u1").String(),
					Key:    "k4",
					Value:  []byte("inserted v4"),
				},
			},
			want1: []*dao.PrivateKVSEtag{
				{
					UserID: faker.UUIDv5("u1").String(),
					Etag:   faker.UUIDv5("e1").String(),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "バケットが空 => バケットは書き込まれず、ETag は削除される",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
				// 削除される
				&dao.PrivateKVSEtag{
					UserID: faker.UUIDv5("u1").String(),
					Etag:   faker.UUIDv5("e1").String(),
				},
			},
			args: args{
				bucket: model.NewPrivateKVSBucket(faker.UUIDv5("u1"), uuid.Nil, nil),
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			var entries []*dao.PrivateKVSEntry
			var etags []*dao.PrivateKVSEtag
			err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				err := r.Save(ctx, tx, tt.args.bucket)
				if err != nil {
					return err
				}

				entries, err = dao.PrivateKVSEntries().All(ctx, tx)
				if err != nil {
					return err
				}
				etags, err = dao.PrivateKVSEtags().All(ctx, tx)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			if diff := cmp.Diff(tt.want, entries, cmpopts.IgnoreFields(dao.PrivateKVSEntry{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
			if diff := cmp.Diff(tt.want1, etags, cmpopts.IgnoreFields(dao.PrivateKVSEtag{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
