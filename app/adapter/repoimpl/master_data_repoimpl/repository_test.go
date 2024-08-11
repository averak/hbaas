package master_data_repoimpl

import (
	"context"
	"testing"
	"time"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/fixture"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Get(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond)

	type args struct {
		revision int
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.MasterData
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "リビジョンが存在する => 取得できる",
			seeds: []fixture.Seed{
				&dao.MasterDatum{
					Revision:  1,
					Content:   []byte("content"),
					IsActive:  true,
					Comment:   "comment",
					CreatedAt: now,
				},
			},
			args: args{
				revision: 1,
			},
			want: model.MasterData{
				Revision:  1,
				Content:   []byte("content"),
				IsActive:  true,
				Comment:   "comment",
				CreatedAt: now,
			},
			wantErr: assert.NoError,
		},
		{
			name:  "リビジョンが存在しない => エラー",
			seeds: []fixture.Seed{},
			args: args{
				revision: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, repository.ErrMasterDataNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)

			conn := testutils.MustDBConn(t)
			var got model.MasterData
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				var err error
				got, err = r.Get(ctx, tx, tt.args.revision)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.EqualExportedValues(t, tt.want, got)
		})
	}
}

func TestRepository_GetActive(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond)

	type args struct {
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.MasterData
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "アクティブなリビジョンが存在する => 取得できる",
			seeds: []fixture.Seed{
				&dao.MasterDatum{
					Revision:  1,
					Content:   []byte("content"),
					IsActive:  true,
					Comment:   "comment",
					CreatedAt: now,
				},
			},
			want: model.MasterData{
				Revision:  1,
				Content:   []byte("content"),
				IsActive:  true,
				Comment:   "comment",
				CreatedAt: now,
			},
			wantErr: assert.NoError,
		},
		{
			name: "アクティブなリビジョンが存在しない => エラー",
			seeds: []fixture.Seed{
				&dao.MasterDatum{
					Revision:  1,
					Content:   []byte("content"),
					IsActive:  false,
					Comment:   "comment",
					CreatedAt: now,
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, repository.ErrActiveMasterDataNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)

			conn := testutils.MustDBConn(t)
			var got model.MasterData
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				var err error
				got, err = r.GetActive(ctx, tx)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.EqualExportedValues(t, tt.want, got)
		})
	}
}

func TestRepository_GetRevisions(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    []int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "リビジョンリストを取得できる",
			seeds: []fixture.Seed{
				&dao.MasterDatum{
					Revision: 1,
				},
				&dao.MasterDatum{
					Revision: 2,
				},
			},
			want:    []int{1, 2},
			wantErr: assert.NoError,
		},
		{
			name:    "リビジョンが存在しない => 空リストを返す",
			seeds:   []fixture.Seed{},
			want:    []int{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)

			conn := testutils.MustDBConn(t)
			var got []int
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				var err error
				got, err = r.GetRevisions(ctx, tx)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRepository_Save(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond)

	type args struct {
		data []model.MasterData
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    []*dao.MasterDatum
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "データを保存できる",
			seeds: []fixture.Seed{
				&dao.MasterDatum{
					Revision: 1,
					IsActive: true,
				},
			},
			args: args{
				data: []model.MasterData{
					// PK が存在する => 更新される
					{
						Revision:  1,
						Content:   []byte("c1"),
						IsActive:  false,
						Comment:   "updated",
						CreatedAt: now,
					},
					// PK が存在しない => 新規作成される
					{
						Revision:  2,
						Content:   []byte("c2"),
						IsActive:  true,
						Comment:   "created",
						CreatedAt: now,
					},
				},
			},
			want: []*dao.MasterDatum{
				{
					Revision:  1,
					Content:   []byte("c1"),
					IsActive:  false,
					Comment:   "updated",
					CreatedAt: now,
				},
				{
					Revision:  2,
					Content:   []byte("c2"),
					IsActive:  true,
					Comment:   "created",
					CreatedAt: now,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "空リスト => 何もしない",
			seeds: []fixture.Seed{
				&dao.MasterDatum{
					Revision:  1,
					IsActive:  true,
					CreatedAt: now,
				},
			},
			args: args{
				data: []model.MasterData{},
			},
			want: []*dao.MasterDatum{
				{
					Revision:  1,
					Content:   []uint8{},
					IsActive:  true,
					CreatedAt: now,
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)

			conn := testutils.MustDBConn(t)
			var dtos []*dao.MasterDatum
			err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				err := r.Save(ctx, tx, tt.args.data...)
				if err != nil {
					return err
				}
				dtos, err = dao.MasterData().All(ctx, tx)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			if diff := cmp.Diff(tt.want, dtos, cmpopts.IgnoreFields(dao.MasterDatum{}, "UpdatedAt")); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
