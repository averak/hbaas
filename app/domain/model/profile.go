package model

import (
	"errors"
	"fmt"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/google/uuid"
)

var (
	ErrUserProfileTooLarge = errors.New("user profile is too large")
)

// プロフィールのデータ構造はプロダクトによって異なり、汎化が難しい。
// そのため、プロフィールをバイナリデータとして扱い、プロダクトごとに独自スキーマを定義させる。
type UserProfile struct {
	UserID uuid.UUID
	raw    []byte
}

func NewUserProfile(userID uuid.UUID, v []byte) (UserProfile, error) {
	if len(v) > numunit.KiB {
		return UserProfile{}, fmt.Errorf("%w: %d bytes", ErrUserProfileTooLarge, len(v))
	}
	return UserProfile{
		UserID: userID,
		raw:    v,
	}, nil
}

func (p UserProfile) Bytes() []byte {
	return p.raw
}
