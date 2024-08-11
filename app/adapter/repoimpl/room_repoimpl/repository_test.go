package room_repoimpl

import (
	"context"
	"testing"

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
	type args struct {
		roomID uuid.UUID
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.Room
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ルームが存在する => 取得できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID:    faker.UUIDv5("u1").String(),
					Email: faker.Email(),
				},
				&dao.User{
					ID:    faker.UUIDv5("u2").String(),
					Email: faker.Email(),
				},
				&dao.User{
					ID:    faker.UUIDv5("u3").String(),
					Email: faker.Email(),
				},
				&dao.Room{
					ID:          faker.UUIDv5("r1").String(),
					OwnerUserID: faker.UUIDv5("u1").String(),
					Type:        int(model.RoomTypePrivate),
					MaxCapacity: 1,
					Secret:      "secret",
					Details:     []byte("details"),
				},
				&dao.RoomUser{
					RoomID: faker.UUIDv5("r1").String(),
					UserID: faker.UUIDv5("u2").String(),
				},
				&dao.RoomUser{
					RoomID: faker.UUIDv5("r1").String(),
					UserID: faker.UUIDv5("u3").String(),
				},
			},
			args: args{
				roomID: faker.UUIDv5("r1"),
			},
			want: model.Room{
				ID:          faker.UUIDv5("r1"),
				OwnerUserID: faker.UUIDv5("u1"),
				Type:        model.RoomTypePrivate,
				MaxCapacity: 1,
				Secret:      "secret",
				Details:     []byte("details"),
				UserIDs:     []uuid.UUID{faker.UUIDv5("u2"), faker.UUIDv5("u3")},
			},
			wantErr: assert.NoError,
		},
		{
			name:  "ルームが存在しない => エラー",
			seeds: []fixture.Seed{},
			args: args{
				roomID: faker.UUIDv5("r1"),
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, repository.ErrRoomNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)

			conn := testutils.MustDBConn(t)
			var got model.Room
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				var err error
				got, err = r.Get(ctx, tx, tt.args.roomID)
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
	type args struct {
		room model.Room
	}
	tests := []struct {
		name    string
		seeds   []fixture.Seed
		args    args
		want    model.Room
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ルームが存在しない => 作成できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID:    faker.UUIDv5("u1").String(),
					Email: faker.Email(),
				},
				&dao.User{
					ID:    faker.UUIDv5("u2").String(),
					Email: faker.Email(),
				},
			},
			args: args{
				room: model.NewRoom(
					faker.UUIDv5("r1"),
					faker.UUIDv5("u1"),
					model.RoomTypePublic,
					1,
					"",
					[]byte("details"),
					[]uuid.UUID{faker.UUIDv5("u1"), faker.UUIDv5("u2")},
				),
			},
			want: model.Room{
				ID:          faker.UUIDv5("r1"),
				OwnerUserID: faker.UUIDv5("u1"),
				Type:        model.RoomTypePublic,
				MaxCapacity: 1,
				Secret:      "",
				Details:     []byte("details"),
				UserIDs:     []uuid.UUID{faker.UUIDv5("u1"), faker.UUIDv5("u2")},
			},
			wantErr: assert.NoError,
		},
		{
			name: "ルームが存在する => 更新できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID:    faker.UUIDv5("u1").String(),
					Email: faker.Email(),
				},
				&dao.User{
					ID:    faker.UUIDv5("u2").String(),
					Email: faker.Email(),
				},
				&dao.User{
					ID:    faker.UUIDv5("u3").String(),
					Email: faker.Email(),
				},
				&dao.Room{
					ID:          faker.UUIDv5("r1").String(),
					OwnerUserID: faker.UUIDv5("u1").String(),
					Type:        int(model.RoomTypePrivate),
					MaxCapacity: 1,
					Secret:      "",
					Details:     []byte("details"),
				},
				&dao.RoomUser{
					RoomID: faker.UUIDv5("r1").String(),
					UserID: faker.UUIDv5("u2").String(),
				},
				&dao.RoomUser{
					RoomID: faker.UUIDv5("r1").String(),
					UserID: faker.UUIDv5("u3").String(),
				},
			},
			args: args{
				room: model.NewRoom(
					faker.UUIDv5("r1"),
					faker.UUIDv5("u1"),
					model.RoomTypePrivate,
					100,
					"secret",
					[]byte("new details"),
					[]uuid.UUID{faker.UUIDv5("u1"), faker.UUIDv5("u2")},
				),
			},
			want: model.Room{
				ID:          faker.UUIDv5("r1"),
				OwnerUserID: faker.UUIDv5("u1"),
				Type:        model.RoomTypePrivate,
				MaxCapacity: 100,
				Secret:      "secret",
				Details:     []byte("new details"),
				UserIDs:     []uuid.UUID{faker.UUIDv5("u1"), faker.UUIDv5("u2")},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)

			conn := testutils.MustDBConn(t)
			var got model.Room
			err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := NewRepository()
				err := r.Save(ctx, tx, tt.args.room)
				if err != nil {
					return err
				}
				got, err = r.Get(ctx, tx, tt.args.room.ID)
				return err
			})
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
