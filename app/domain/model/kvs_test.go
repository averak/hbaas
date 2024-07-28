package model

import (
	"fmt"
	"testing"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKVSBucket_HasKey(t *testing.T) {
	type fields struct {
		raw []KVSEntry
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "キーが存在する => true",
			fields: fields{
				raw: []KVSEntry{
					{
						Key: "k1",
					},
				},
			},
			args: args{
				key: "k1",
			},
			want: true,
		},
		{
			name: "キーが存在しない => false",
			fields: fields{
				raw: []KVSEntry{
					{
						Key: "k1",
					},
				},
			},
			args: args{
				key: "k2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := KVSBucket{
				raw: tt.fields.raw,
			}
			assert.Equalf(t, tt.want, m.HasKey(tt.args.key), "HasKey(%v)", tt.args.key)
		})
	}
}

func TestKVSBucket_Set(t *testing.T) {
	type fields struct {
		raw []KVSEntry
	}
	type args struct {
		entries []KVSEntry
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   KVSBucket
	}{
		{
			name: "キーが存在する場合は更新、存在しない場合は追加される",
			fields: fields{
				raw: []KVSEntry{
					{
						Key:   "k1",
						Value: []byte("v1"),
					},
					{
						Key:   "k2",
						Value: []byte("v2"),
					},
				},
			},
			args: args{
				entries: []KVSEntry{
					{
						Key:   "k2",
						Value: []byte("updated v2"),
					},
					{
						Key:   "k3",
						Value: []byte("inserted v3"),
					},
				},
			},
			want: KVSBucket{
				raw: []KVSEntry{
					{
						Key:   "k1",
						Value: []byte("v1"),
					},
					{
						Key:   "k2",
						Value: []byte("updated v2"),
					},
					{
						Key:   "k3",
						Value: []byte("inserted v3"),
					},
				},
			},
		},
		{
			name: "空リスト => 何もしない",
			fields: fields{
				raw: []KVSEntry{
					{
						Key:   "k1",
						Value: []byte("v1"),
					},
				},
			},
			args: args{
				entries: []KVSEntry{},
			},
			want: KVSBucket{
				raw: []KVSEntry{
					{
						Key:   "k1",
						Value: []byte("v1"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := KVSBucket{
				raw: tt.fields.raw,
			}
			m.Set(tt.args.entries)
			if diff := cmp.Diff(tt.want, m, cmp.AllowUnexported(KVSBucket{})); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}

func TestNewKVSEntry(t *testing.T) {
	type args struct {
		key   string
		value []byte
	}
	tests := []struct {
		name    string
		args    args
		want    KVSEntry
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "バイナリサイズ <= 100KiB の場合 => 成功",
			args: args{
				key:   "k1",
				value: make([]byte, 100*numunit.KiB),
			},
			want: KVSEntry{
				Key:   "k1",
				Value: make([]byte, 100*numunit.KiB),
			},
			wantErr: assert.NoError,
		},
		{
			name: "バイナリサイズ > 100KiB の場合 => エラー",
			args: args{
				key:   "k1",
				value: make([]byte, 100*numunit.KiB+1),
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrKVSEntryValueTooLarge)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKVSEntry(tt.args.key, tt.args.value)
			if !tt.wantErr(t, err, fmt.Sprintf("NewKVSEntry(%v, %v)", tt.args.key, tt.args.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewKVSEntry(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestKVSEntry_IsEmpty(t *testing.T) {
	type fields struct {
		Key   string
		Value []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "バイナリサイズが 0 の場合 => true",
			fields: fields{
				Value: []byte{},
			},
			want: true,
		},
		{
			name: "バイナリサイズが 0 でない場合 => false",
			fields: fields{
				Value: []byte{0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := KVSEntry{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, e.IsEmpty(), "IsEmpty()")
		})
	}
}

func TestKVSCriteria_IsEmpty(t *testing.T) {
	type fields struct {
		ExactMatch  []string
		PrefixMatch []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ExactMatch が空 && PrefixMatch が空 => false",
			fields: fields{
				ExactMatch:  []string{},
				PrefixMatch: []string{},
			},
			want: true,
		},
		{
			name: "ExactMatch が空 && PrefixMatch が空でない => false",
			fields: fields{
				ExactMatch:  []string{},
				PrefixMatch: []string{"k1"},
			},
			want: false,
		},
		{
			name: "ExactMatch が空でない && PrefixMatch が空 => false",
			fields: fields{
				ExactMatch:  []string{"k1"},
				PrefixMatch: []string{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := KVSCriteria{
				ExactMatch:  tt.fields.ExactMatch,
				PrefixMatch: tt.fields.PrefixMatch,
			}
			assert.Equalf(t, tt.want, c.IsEmpty(), "IsEmpty()")
		})
	}
}

func TestPrivateKVSBucket_Set(t *testing.T) {
	type fields struct {
		KVSBucket KVSBucket
		etag      uuid.UUID
	}
	type args struct {
		etag    uuid.UUID
		entries []KVSEntry
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		then   func(t *testing.T, m PrivateKVSBucket, err error)
	}{
		{
			name: "etag が一致する場合 => エントリをセットし、etag を更新する",
			fields: fields{
				KVSBucket: KVSBucket{
					raw: []KVSEntry{
						{
							Key:   "k1",
							Value: []byte("v1"),
						},
						{
							Key:   "k2",
							Value: []byte("v2"),
						},
					},
				},
				etag: faker.UUIDv5("e1"),
			},
			args: args{
				etag: faker.UUIDv5("e1"),
				entries: []KVSEntry{
					{
						Key:   "k2",
						Value: []byte("updated v2"),
					},
					{
						Key:   "k3",
						Value: []byte("inserted v3"),
					},
				},
			},
			then: func(t *testing.T, m PrivateKVSBucket, err error) {
				require.NoError(t, err)

				want := KVSBucket{
					raw: []KVSEntry{
						{
							Key:   "k1",
							Value: []byte("v1"),
						},
						{
							Key:   "k2",
							Value: []byte("updated v2"),
						},
						{
							Key:   "k3",
							Value: []byte("inserted v3"),
						},
					},
				}
				assert.Equal(t, want, m.KVSBucket)
				assert.NotEqual(t, faker.UUIDv5("e1"), m.etag)
			},
		},
		{
			name: "etag が一致しない場合 => エラー",
			fields: fields{
				KVSBucket: KVSBucket{},
				etag:      faker.UUIDv5("e1"),
			},
			args: args{
				etag:    faker.UUIDv5("e2"),
				entries: []KVSEntry{},
			},
			then: func(t *testing.T, m PrivateKVSBucket, err error) {
				require.ErrorIs(t, err, ErrPrivateKVSETagMismatch)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := PrivateKVSBucket{
				KVSBucket: tt.fields.KVSBucket,
				etag:      tt.fields.etag,
			}
			err := m.Set(tt.args.etag, tt.args.entries)
			tt.then(t, m, err)
		})
	}
}
