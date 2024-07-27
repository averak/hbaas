package model

import (
	"testing"

	"github.com/averak/hbaas/testutils/faker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserProfile(t *testing.T) {
	type args struct {
		userID uuid.UUID
		v      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    UserProfile
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "バイナリサイズ <= 1KiB の場合 => 成功",
			args: args{
				userID: faker.UUIDv5("u1"),
				v:      make([]byte, 1024),
			},
			want: UserProfile{
				UserID: faker.UUIDv5("u1"),
				raw:    make([]byte, 1024),
			},
			wantErr: assert.NoError,
		},
		{
			name: "バイナリサイズ > 1KiB の場合 => エラー",
			args: args{
				userID: faker.UUIDv5("u1"),
				v:      make([]byte, 1025),
			},
			want: UserProfile{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrUserProfileTooLarge)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserProfile(tt.args.userID, tt.args.v)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
