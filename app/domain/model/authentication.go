package model

import (
	"time"

	"github.com/google/uuid"
)

// UserAuthentication はユーザの認証に必要な情報を提供します。
// 認証/再認証のみで利用され、セッション管理のコンテキストでは利用されません。
type UserAuthentication struct {
	UserID uuid.UUID
	// BaaS ごとにユーザIDの表現が異なるため、文字列型にしている。
	BaasUserID          string
	LastAuthenticatedAt time.Time
}

func NewUserAuthentication(userID uuid.UUID, baasUserID string, lastAuthenticatedAt time.Time) UserAuthentication {
	return UserAuthentication{
		UserID:              userID,
		BaasUserID:          baasUserID,
		LastAuthenticatedAt: lastAuthenticatedAt,
	}
}

func (a *UserAuthentication) Refresh(now time.Time) {
	a.LastAuthenticatedAt = now
}
