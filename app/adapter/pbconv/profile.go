package pbconv

import (
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/protobuf/resource"
)

func ToProfilePb(v model.UserProfile) *resource.UserProfile {
	return &resource.UserProfile{
		UserId: v.UserID.String(),
		Data:   v.Bytes(),
	}
}
