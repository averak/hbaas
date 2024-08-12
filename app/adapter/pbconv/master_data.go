package pbconv

import (
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/protobuf/resource"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToMasterDataPb(m model.MasterData) *resource.MasterData {
	return &resource.MasterData{
		Revision:  int64(m.Revision),
		Content:   m.Content,
		IsActive:  m.IsActive,
		Comment:   m.Comment,
		CreatedAt: timestamppb.New(m.CreatedAt),
	}
}
