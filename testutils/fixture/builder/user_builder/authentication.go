package user_builder

import (
	"time"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/google/uuid"
)

type UserAuthenticationBuilder struct {
	data model.UserAuthentication
}

func NewUserAuthentication(userID uuid.UUID) *UserAuthenticationBuilder {
	return &UserAuthenticationBuilder{
		data: model.NewUserAuthentication(userID, userID.String(), time.Now()),
	}
}

func (b UserAuthenticationBuilder) Build() model.UserAuthentication {
	return b.data
}
