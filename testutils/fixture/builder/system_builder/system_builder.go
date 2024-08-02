package system_builder

import "github.com/averak/hbaas/app/domain/model"

type Data struct {
	GlobalKVSBucket *model.GlobalKVSBucket
	LeaderBoard     []*model.LeaderBoard
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

func (b *SystemBuilder) LeaderBoard(v model.LeaderBoard) *SystemBuilder {
	b.data.LeaderBoard = append(b.data.LeaderBoard, &v)
	return b
}
