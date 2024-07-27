package system_builder

import "github.com/averak/hbaas/app/domain/model"

type Data struct {
	GlobalKVSBucket *model.GlobalKVSBucket
}

type SystemBuilder struct {
	data Data
}

func New() *SystemBuilder {
	return &SystemBuilder{}
}

func (b SystemBuilder) Build() Data {
	return b.data
}

func (b *SystemBuilder) GlobalKVSBucket(v model.GlobalKVSBucket) *SystemBuilder {
	b.data.GlobalKVSBucket = &v
	return b
}
