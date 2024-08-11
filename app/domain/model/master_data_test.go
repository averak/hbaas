package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/stretchr/testify/assert"
)

func TestNewMasterData(t *testing.T) {
	type args struct {
		revision  int
		content   []byte
		isActive  bool
		comment   string
		createdAt time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    MasterData
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "バイナリサイズ <= 1MiB の場合 => 成功",
			args: args{
				content: make([]byte, numunit.MiB),
			},
			want: MasterData{
				Content: make([]byte, numunit.MiB),
			},
			wantErr: assert.NoError,
		},
		{
			name: "バイナリサイズ > 1MiB の場合 => エラー",
			args: args{
				content: make([]byte, numunit.MiB+1),
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMasterData(tt.args.revision, tt.args.content, tt.args.isActive, tt.args.comment, tt.args.createdAt)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewMasterDataService(t *testing.T) {
	type args struct {
		active    MasterData
		revisions []int
	}
	tests := []struct {
		name    string
		args    args
		want    *MasterDataService
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "active がアクティブの場合 => 成功",
			args: args{
				active: MasterData{
					IsActive: true,
				},
				revisions: []int{3, 2, 1},
			},
			want: &MasterDataService{
				active: MasterData{
					IsActive: true,
				},
				// ソートされる
				sortedRevisions: []int{1, 2, 3},
			},
			wantErr: assert.NoError,
		},
		{
			name: "active が非アクティブの場合 => エラー",
			args: args{
				active: MasterData{
					IsActive: false,
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "revisions が空の場合 => エラー",
			args: args{
				active: MasterData{
					IsActive: true,
				},
				revisions: []int{},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMasterDataService(tt.args.active, tt.args.revisions)
			if !tt.wantErr(t, err, fmt.Sprintf("NewMasterDataService(%v, %v)", tt.args.active, tt.args.revisions)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewMasterDataService(%v, %v)", tt.args.active, tt.args.revisions)
		})
	}
}

func TestMasterDataService_SwitchActive(t *testing.T) {
	type fields struct {
		active MasterData
	}
	type args struct {
		target MasterData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MasterData
		want1   MasterData
		want2   *MasterDataService
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "target が非アクティブの場合 => 成功",
			fields: fields{
				active: MasterData{
					Revision: 1,
					IsActive: true,
					Content:  []byte("m1"),
				},
			},
			args: args{
				target: MasterData{
					Revision: 2,
					IsActive: false,
					Content:  []byte("m2"),
				},
			},
			want: MasterData{
				Revision: 1,
				IsActive: false,
				Content:  []byte("m1"),
			},
			want1: MasterData{
				Revision: 2,
				IsActive: true,
				Content:  []byte("m2"),
			},
			want2: &MasterDataService{
				active: MasterData{
					Revision: 2,
					IsActive: true,
					Content:  []byte("m2"),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "target がアクティブの場合 => エラー",
			fields: fields{
				active: MasterData{
					Revision: 1,
					IsActive: true,
				},
			},
			args: args{
				target: MasterData{
					Revision: 2,
					IsActive: true,
				},
			},
			want2: &MasterDataService{
				active: MasterData{
					Revision: 1,
					IsActive: true,
				},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MasterDataService{
				active: tt.fields.active,
			}
			got, got1, err := s.SwitchActive(tt.args.target)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, s)
		})
	}
}

func TestMasterDataService_NewNextRevision(t *testing.T) {
	type fields struct {
		revisions []int
	}
	type args struct {
		content []byte
		comment string
		now     time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MasterData
		want1   *MasterDataService
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "新規リビジョンを作成できる",
			fields: fields{
				revisions: []int{1, 2, 3},
			},
			args: args{
				content: []byte("m4"),
			},
			want: MasterData{
				Revision: 4,
				Content:  []byte("m4"),
			},
			want1: &MasterDataService{
				sortedRevisions: []int{1, 2, 3, 4},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MasterDataService{
				sortedRevisions: tt.fields.revisions,
			}
			got, err := s.NewNextRevision(tt.args.content, tt.args.comment, tt.args.now)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, s)
		})
	}
}
