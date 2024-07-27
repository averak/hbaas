package model

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserDeactivated = errors.New("user is deactivated")
)

type UserStatus int

const (
	// 設定が未完了で、他ユーザからの操作を受け付けない状態。
	UserStatusPending UserStatus = 1 + iota
	// アカウントが利用可能な状態。
	UserStatusActive
	// アカウントが削除され、利用不可な状態。
	// この状態になると、二度と利用できなくなる。
	UserStatusDeactivated
)

type User struct {
	ID     uuid.UUID
	Email  string
	Status UserStatus
}

func NewUser(id uuid.UUID, email string, status UserStatus) User {
	return User{
		ID:     id,
		Email:  email,
		Status: status,
	}
}

func (u *User) Activate() error {
	if u.Status == UserStatusDeactivated {
		return ErrUserDeactivated
	}
	u.Status = UserStatusActive
	return nil
}

func (u User) IsUnavailable() bool {
	return u.Status == UserStatusDeactivated
}
