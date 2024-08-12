package system_builder

import (
	"log"
	"time"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/averak/hbaas/app/domain/model"
)

type MasterDataBuilder struct {
	data model.MasterData
}

func NewMasterDataBuilder(revision int) *MasterDataBuilder {
	data, err := model.NewMasterData(revision, make([]byte, 10*numunit.B), false, "comment", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	return &MasterDataBuilder{
		data: data,
	}
}

func (b MasterDataBuilder) Build() model.MasterData {
	return b.data
}

func (b *MasterDataBuilder) Content(v []byte) *MasterDataBuilder {
	b.data.Content = v
	return b
}

func (b *MasterDataBuilder) IsActive(v bool) *MasterDataBuilder {
	b.data.IsActive = v
	return b
}

func (b *MasterDataBuilder) Comment(v string) *MasterDataBuilder {
	b.data.Comment = v
	return b
}

func (b *MasterDataBuilder) CreatedAt(v time.Time) *MasterDataBuilder {
	b.data.CreatedAt = v
	return b
}
