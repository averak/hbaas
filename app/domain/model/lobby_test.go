package model

import (
	"fmt"
	"testing"

	"github.com/averak/hbaas/testutils/faker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLobbyService_CreatePrivateRoom(t *testing.T) {
	type args struct {
		id          uuid.UUID
		ownerUserID uuid.UUID
		maxCapacity int
		secret      string
		detail      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Room
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "合言葉が空文字列ではない => 作成できる",
			args: args{
				id:          faker.UUIDv5("r1"),
				ownerUserID: faker.UUIDv5("u1"),
				maxCapacity: 10,
				secret:      "secret",
				detail:      []byte("detail"),
			},
			want: Room{
				ID:          faker.UUIDv5("r1"),
				OwnerUserID: faker.UUIDv5("u1"),
				Type:        RoomTypePrivate,
				MaxCapacity: 10,
				Secret:      "secret",
				Details:     []byte("detail"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "合言葉が空文字列 => エラー",
			args: args{
				secret: "",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrRoomSecretEmpty)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LobbyService{}
			got, err := s.CreatePrivateRoom(tt.args.id, tt.args.ownerUserID, tt.args.maxCapacity, tt.args.secret, tt.args.detail)
			if !tt.wantErr(t, err, fmt.Sprintf("CreatePrivateRoom(%v, %v, %v, %v, %v)", tt.args.id, tt.args.ownerUserID, tt.args.maxCapacity, tt.args.secret, tt.args.detail)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CreatePrivateRoom(%v, %v, %v, %v, %v)", tt.args.id, tt.args.ownerUserID, tt.args.maxCapacity, tt.args.secret, tt.args.detail)
		})
	}
}
