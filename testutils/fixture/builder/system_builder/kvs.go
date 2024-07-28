package system_builder

import (
	"github.com/averak/hbaas/app/domain/model"
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
