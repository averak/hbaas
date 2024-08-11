package user_builder

import (
	"log"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/google/uuid"
)

type UserProfileBuilder struct {
	data model.UserProfile
}

func NewUserProfile(userID uuid.UUID) *UserProfileBuilder {
	data, err := model.NewUserProfile(userID, make([]byte, numunit.B))
	if err != nil {
		log.Fatal(err)
	}
	return &UserProfileBuilder{
		data: data,
	}
}

func (b UserProfileBuilder) Build() model.UserProfile {
	return b.data
}

func (b *UserProfileBuilder) Raw(v []byte) *UserProfileBuilder {
	data, err := model.NewUserProfile(b.data.UserID, v)
	if err != nil {
		log.Fatal(err)
	}
	b.data = data
	return b
}
