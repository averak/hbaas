package leader_board_repoimpl

import (
	"context"
	"testing"
	"time"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/averak/hbaas/testutils/fixture"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Get(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.LeaderBoard
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "リーダーボードが存在する => 取得できる",
			seeds: []fixture.Seed{
				&dao.LeaderBoard{
					ID: faker.UUIDv5("l1").String(),
				},
				&dao.LeaderBoard{
					ID: faker.UUIDv5("l2").String(),
				},
				&dao.LeaderBoardScore{
					LeaderBoardID: faker.UUIDv5("l1").String(),
					ScoreID:       faker.UUIDv5("s1").String(),
					Score:         1,
					Timestamp:     now,
				},
				&dao.LeaderBoardScore{
					LeaderBoardID: faker.UUIDv5("l1").String(),
					ScoreID:       faker.UUIDv5("s2").String(),
					Score:         2,
					Timestamp:     now,
				},
				&dao.LeaderBoardScore{
					LeaderBoardID: faker.UUIDv5("l2").String(),
					ScoreID:       faker.UUIDv5("s3").String(),
					Score:         3,
					Timestamp:     now,
				},
			},
			args: args{
				id: faker.UUIDv5("l1").String(),
			},
			want: model.LeaderBoard{
				ID: faker.UUIDv5("l1").String(),
				Scores: []model.LeaderBoardScore{
					{
						ScoreID:   faker.UUIDv5("s2").String(),
						Score:     2,
						Timestamp: now,
					},
					{
						ScoreID:   faker.UUIDv5("s1").String(),
						Score:     1,
						Timestamp: now,
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name:  "リーダーボードが存在しない => 空のリーダーボードを返す",
			seeds: []fixture.Seed{},
			args: args{
				id: faker.UUIDv5("not exists").String(),
			},
			want: model.LeaderBoard{
				ID:     faker.UUIDv5("not exists").String(),
				Scores: nil,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)

			conn := testutils.MustDBConn(t)
			var got model.LeaderBoard
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				var err error
				got, err = r.Get(ctx, tx, tt.args.id)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.EqualExportedValues(t, tt.want, got)
		})
	}
}

func TestRepository_Save(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond)

	type args struct {
		leaderBoard model.LeaderBoard
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.LeaderBoard
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "リーダーボードが存在する => 更新できる",
			seeds: []fixture.Seed{
				&dao.LeaderBoard{
					ID: faker.UUIDv5("l1").String(),
				},
				&dao.LeaderBoardScore{
					LeaderBoardID: faker.UUIDv5("l1").String(),
					ScoreID:       faker.UUIDv5("s1").String(),
					Score:         1,
					Timestamp:     now,
				},
				&dao.LeaderBoardScore{
					LeaderBoardID: faker.UUIDv5("l1").String(),
					ScoreID:       faker.UUIDv5("s2").String(),
					Score:         2,
					Timestamp:     now,
				},
			},
			args: args{
				leaderBoard: model.NewLeaderBoard(
					faker.UUIDv5("l1").String(),
					[]model.LeaderBoardScore{
						model.NewLeaderBoardScore(faker.UUIDv5("s2").String(), 10, now),
						model.NewLeaderBoardScore(faker.UUIDv5("s3").String(), 20, now),
					},
				),
			},
			want: model.LeaderBoard{
				ID: faker.UUIDv5("l1").String(),
				Scores: []model.LeaderBoardScore{
					model.NewLeaderBoardScore(faker.UUIDv5("s3").String(), 20, now),
					model.NewLeaderBoardScore(faker.UUIDv5("s2").String(), 10, now),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name:  "リーダーボードが存在しない => 作成できる",
			seeds: []fixture.Seed{},
			args: args{
				leaderBoard: model.NewLeaderBoard(
					faker.UUIDv5("l1").String(),
					[]model.LeaderBoardScore{
						model.NewLeaderBoardScore(faker.UUIDv5("s1").String(), 1, now),
						model.NewLeaderBoardScore(faker.UUIDv5("s2").String(), 2, now),
					},
				),
			},
			want: model.LeaderBoard{
				ID: faker.UUIDv5("l1").String(),
				Scores: []model.LeaderBoardScore{
					model.NewLeaderBoardScore(faker.UUIDv5("s2").String(), 2, now),
					model.NewLeaderBoardScore(faker.UUIDv5("s1").String(), 1, now),
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
			var got model.LeaderBoard
			err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				err := r.Save(ctx, tx, tt.args.leaderBoard)
				if err != nil {
					return err
				}
				got, err = r.Get(ctx, tx, tt.args.leaderBoard.ID)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.EqualExportedValues(t, tt.want, got)
		})
	}
}
