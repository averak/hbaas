package system_builder

import (
	"testing"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/google/uuid"
)

type GlobalKVSBucketBuilder struct {
	data model.GlobalKVSBucket
}

func NewGlobalKVSBucketBuilder() *GlobalKVSBucketBuilder {
	return &GlobalKVSBucketBuilder{
		data: model.NewGlobalKVSBucket(nil),
	}
}

func (b GlobalKVSBucketBuilder) Build() model.GlobalKVSBucket {
	return b.data
}

func (b *GlobalKVSBucketBuilder) Entries(v ...model.KVSEntry) *GlobalKVSBucketBuilder {
	b.data.Set(v...)
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
