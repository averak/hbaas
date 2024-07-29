package user_builder

import (
	"fmt"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/google/uuid"
)

type Data struct {
	User             model.User
	Authentication   *model.UserAuthentication
	PrivateKVSBucket *model.PrivateKVSBucket
}

type UserBuilder struct {
	data *Data
}

func New(userID uuid.UUID) *UserBuilder {
	return &UserBuilder{
		data: &Data{
			User: model.NewUser(userID, fmt.Sprintf("%s@example.com", userID), model.UserStatusActive),
		},
	}
}

func (b UserBuilder) Build() Data {
	return *b.data
}

func (b *UserBuilder) Status(v model.UserStatus) *UserBuilder {
	b.data.User.Status = v
	return b
}

func (b *UserBuilder) Authentication(v model.UserAuthentication) *UserBuilder {
	b.data.Authentication = &v
	return b
}

func (b *UserBuilder) PrivateKVSBucket(v model.PrivateKVSBucket) *UserBuilder {
	b.data.PrivateKVSBucket = &v
	return b
}
