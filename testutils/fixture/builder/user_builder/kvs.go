package user_builder

import (
	"log"
	"testing"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/google/uuid"
)

type PrivateKVSBucketBuilder struct {
	data model.PrivateKVSBucket
}

func NewPrivateKVSBucketBuilder(userID uuid.UUID) *PrivateKVSBucketBuilder {
	return &PrivateKVSBucketBuilder{
		data: model.NewPrivateKVSBucket(userID, uuid.New(), nil),
	}
}

func (b PrivateKVSBucketBuilder) Build() model.PrivateKVSBucket {
	return b.data
}

func (b *PrivateKVSBucketBuilder) Entries(v ...model.KVSEntry) *PrivateKVSBucketBuilder {
	err := b.data.Set(b.data.ETag(), v)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func (b *PrivateKVSBucketBuilder) ETag(v uuid.UUID) *PrivateKVSBucketBuilder {
	b.data = model.NewPrivateKVSBucket(b.data.UserID, v, b.data.Raw())
	return b
}

type KVSEntryBuilder struct {
	data model.KVSEntry
}

func NewKVSEntryBuilder(t *testing.T) *KVSEntryBuilder {
	v, err := model.NewKVSEntry(uuid.New().String(), make([]byte, 1*numunit.B))
	if err != nil {
		t.Fatal(err)
	}
	return &KVSEntryBuilder{
		data: v,
	}
}

func (b KVSEntryBuilder) Build() model.KVSEntry {
	return b.data
}

func (b *KVSEntryBuilder) Key(v string) *KVSEntryBuilder {
	b.data.Key = v
	return b
}

func (b *KVSEntryBuilder) Value(v []byte) *KVSEntryBuilder {
	b.data.Value = v
	return b
}
