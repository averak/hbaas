package model

import (
	"testing"

	"github.com/averak/hbaas/testutils/faker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRoom_Join(t *testing.T) {
	type fields struct {
		Type        RoomType
		MaxCapacity int
		Secret      string
		UserIDs     []uuid.UUID
	}
	type args struct {
		userID uuid.UUID
		secret string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Room
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "[パブリックルーム] 入室できる",
			fields: fields{
				Type:        RoomTypePublic,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs:     []uuid.UUID{},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
				secret: "secret",
			},
			want: &Room{
				Type:        RoomTypePublic,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs:     []uuid.UUID{faker.UUIDv5("u1")},
			},
			wantErr: assert.NoError,
		},
		{
			name: "[バブリックルーム] 入室済みの場合 => エラー",
			fields: fields{
				Type:        RoomTypePublic,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
				},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
				secret: "secret",
			},
			want: &Room{
				Type:        RoomTypePublic,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs:     []uuid.UUID{faker.UUIDv5("u1")},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, ErrUserAlreadyInRoom, err)
			},
		},
		{
			name: "[パブリックルーム] 既に最大人数の場合 => エラー",
			fields: fields{
				Type:        RoomTypePublic,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
					faker.UUIDv5("u2"),
				},
			},
			args: args{
				userID: faker.UUIDv5("u3"),
				secret: "secret",
			},
			want: &Room{
				Type:        RoomTypePublic,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
					faker.UUIDv5("u2"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, ErrRoomMaxCapacity, err)
			},
		},
		{
			name: "[プレイベートルーム] 合言葉が一致した場合 => 入室できる",
			fields: fields{
				Type:        RoomTypePrivate,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs:     []uuid.UUID{},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
				secret: "secret",
			},
			want: &Room{
				Type:        RoomTypePrivate,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs:     []uuid.UUID{faker.UUIDv5("u1")},
			},
			wantErr: assert.NoError,
		},
		{
			name: "[プレイベートルーム] 合言葉が不一致の場合 => エラー",
			fields: fields{
				Type:        RoomTypePrivate,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs:     []uuid.UUID{},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
				secret: "invalid",
			},
			want: &Room{
				Type:        RoomTypePrivate,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs:     []uuid.UUID{},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, ErrRoomSecretMismatch, err)
			},
		},
		{
			name: "[プレイベートルーム] 入室済みの場合 => エラー",
			fields: fields{
				Type:        RoomTypePrivate,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
				},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
				secret: "secret",
			},
			want: &Room{
				Type:        RoomTypePrivate,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs:     []uuid.UUID{faker.UUIDv5("u1")},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, ErrUserAlreadyInRoom, err)
			},
		},
		{
			name: "[プレイベートルーム] 既に最大人数の場合 => エラー",
			fields: fields{
				Type:        RoomTypePrivate,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
					faker.UUIDv5("u2"),
				},
			},
			args: args{
				userID: faker.UUIDv5("u3"),
				secret: "secret",
			},
			want: &Room{
				Type:        RoomTypePrivate,
				MaxCapacity: 2,
				Secret:      "secret",
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
					faker.UUIDv5("u2"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, ErrRoomMaxCapacity, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				Type:        tt.fields.Type,
				MaxCapacity: tt.fields.MaxCapacity,
				Secret:      tt.fields.Secret,
				UserIDs:     tt.fields.UserIDs,
			}
			err := r.Join(tt.args.userID, tt.args.secret)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, r)
		})
	}
}

func TestRoom_Leave(t *testing.T) {
	type fields struct {
		UserIDs []uuid.UUID
	}
	type args struct {
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Room
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ルームに参加している場合 => 退室できる",
			fields: fields{
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
				},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
			},
			want: &Room{
				UserIDs: []uuid.UUID{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "ルームに参加していない場合 => エラー",
			fields: fields{
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
				},
			},
			args: args{
				userID: faker.UUIDv5("u2"),
			},
			want: &Room{
				UserIDs: []uuid.UUID{
					faker.UUIDv5("u1"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, ErrUserNotInRoom, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				UserIDs: tt.fields.UserIDs,
			}
			if !tt.wantErr(t, r.Leave(tt.args.userID)) {
				return
			}
			assert.Equal(t, tt.want, r)
		})
	}
}
