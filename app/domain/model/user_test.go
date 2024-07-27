package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Activate(t *testing.T) {
	type fields struct {
		Status UserStatus
	}
	tests := []struct {
		name    string
		fields  fields
		want    User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Pending の場合 => アクティベートできる",
			fields: fields{
				Status: UserStatusPending,
			},
			want: User{
				Status: UserStatusActive,
			},
			wantErr: assert.NoError,
		},
		{
			// 冪等性を保証するために、エラーにしない。
			name: "Active の場合 => アクティベートできる",
			fields: fields{
				Status: UserStatusActive,
			},
			want: User{
				Status: UserStatusActive,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Deactivated の場合 => エラー",
			fields: fields{
				Status: UserStatusDeactivated,
			},
			want: User{
				Status: UserStatusDeactivated,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrUserDeactivated)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Status: tt.fields.Status,
			}
			err := u.Activate()
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, u)
		})
	}
}
