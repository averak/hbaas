package model

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserDeactivated = errors.New("user is deactivated")
)

const (
	// DeletedUserEmail は、削除されたユーザのメールアドレスを表します。
	// メールアドレスの UNIQUE 制約は論理削除されていないレコードのみを対象としているため、削除済みユーザのメールアドレスが重複しても問題ない。
	DeletedUserEmail = ""
)

type UserStatus int

const (
	// 設定が未完了で、他ユーザには公開されていない状態。
	// サインイン後、設定を完了するまでこの状態が続く。
	UserStatusPending UserStatus = 1 + iota
	// アカウントが利用可能で、他ユーザにも公開されている状態。
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

func (u *User) Delete() {
	u.Status = UserStatusDeactivated
	u.Email = DeletedUserEmail
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
