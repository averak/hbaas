package global_kvs_repoimpl

import (
	"context"
	"testing"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
	"github.com/averak/hbaas/testutils/fixture"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_Get(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type given struct {
		seeds []fixture.Seed
	}
	type when struct {
		criteria model.KVSCriteria
	}
	type then = func(t *testing.T, got model.GlobalKVSBucket, err error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "レコードが存在する状態で",
			Given: given{
				seeds: []fixture.Seed{
					&dao.GlobalKVSEntry{
						Key:   "group1:key1",
						Value: []byte("v1"),
					},
					&dao.GlobalKVSEntry{
						Key:   "group1:key2",
						Value: []byte("v2"),
					},
					&dao.GlobalKVSEntry{
						Key:   "group2:key1",
						Value: []byte("v3"),
					},
					&dao.GlobalKVSEntry{
						Key:   "group2:key2",
						Value: []byte("v4"),
					},
					&dao.GlobalKVSEntry{
						Key:   "group3:key1",
						Value: []byte("v5"),
					},
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "前方一致で取得できる",
					When: when{
						criteria: model.KVSCriteria{
							PrefixMatch: []string{"group1"},
						},
					},
					Then: func(t *testing.T, got model.GlobalKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewGlobalKVSBucket(
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
						criteria: model.KVSCriteria{
							ExactMatch: []string{
								"group1:key1",
								"group2:key1",
							},
						},
					},
					Then: func(t *testing.T, got model.GlobalKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewGlobalKVSBucket(
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
						criteria: model.KVSCriteria{
							ExactMatch:  []string{"group1:key1"},
							PrefixMatch: []string{"group2"},
						},
					},
					Then: func(t *testing.T, got model.GlobalKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewGlobalKVSBucket(
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
						criteria: model.KVSCriteria{},
					},
					Then: func(t *testing.T, got model.GlobalKVSBucket, err error) {
						require.NoError(t, err)

						want := model.NewGlobalKVSBucket(nil)
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

			var got model.GlobalKVSBucket
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				var err error
				got, err = r.Get(ctx, tx, when.criteria)
				return err
			})
			then(t, got, err)
		})
	}
}

func TestRepository_Save(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type args struct {
		bucket model.GlobalKVSBucket
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    []*dao.GlobalKVSEntry
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "エントリを保存できる",
			seeds: []fixture.Seed{
				&dao.GlobalKVSEntry{
					Key:   "k1",
					Value: []byte("v1"),
				},
				&dao.GlobalKVSEntry{
					Key:   "k2",
					Value: []byte("v2"),
				},
				&dao.GlobalKVSEntry{
					Key:   "k3",
					Value: []byte("v3"),
				},
			},
			args: args{
				bucket: model.NewGlobalKVSBucket(
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
			want: []*dao.GlobalKVSEntry{
				{
					Key:   "k1",
					Value: []byte("v1"),
				},
				{
					Key:   "k2",
					Value: []byte("updated v2"),
				},
				{
					Key:   "k4",
					Value: []byte("inserted v4"),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name:  "バケットが空 => 何もしない",
			seeds: []fixture.Seed{},
			args: args{
				bucket: model.NewGlobalKVSBucket(nil),
			},
			want:    nil,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			var dtos []*dao.GlobalKVSEntry
			err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				err := r.Save(ctx, tx, tt.args.bucket)
				if err != nil {
					return err
				}

				dtos, err = dao.GlobalKVSEntries().All(ctx, tx)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			if diff := cmp.Diff(tt.want, dtos, cmpopts.IgnoreFields(dao.GlobalKVSEntry{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
