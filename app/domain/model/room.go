package model

import (
	"errors"
	"slices"

	"github.com/averak/hbaas/pkg/vector"
	"github.com/google/uuid"
)

var (
	ErrRoomSecretEmpty    = errors.New("room secret is empty")
	ErrRoomSecretMismatch = errors.New("room secret mismatch")
	ErrRoomMaxCapacity    = errors.New("room is full")
	ErrUserAlreadyInRoom  = errors.New("user is already in room")
	ErrUserNotInRoom      = errors.New("user is not in room")
)

type RoomType int

const (
	// 誰でも自由に出入りできるルーム。
	RoomTypePublic RoomType = iota + 1
	// 合言葉を知っているユーザのみが出入りできるルーム。
	RoomTypePrivate
)

// Room は、WebRTC のルームを表します。
type Room struct {
	ID          uuid.UUID
	OwnerUserID uuid.UUID
	Type        RoomType
	MaxCapacity int
	Secret      string
	Details     []byte

	UserIDs []uuid.UUID
}

func NewRoom(id uuid.UUID, ownerUserID uuid.UUID, roomType RoomType, maxCapacity int, secret string, details []byte, userIDs []uuid.UUID) Room {
	return Room{
		ID:          id,
		OwnerUserID: ownerUserID,
		Type:        roomType,
		MaxCapacity: maxCapacity,
		Secret:      secret,
		Details:     details,
		UserIDs:     userIDs,
	}
}

func (r *Room) Join(userID uuid.UUID, secret string) error {
	if r.Type == RoomTypePrivate && r.Secret != secret {
		return ErrRoomSecretMismatch
	}
	if len(r.UserIDs) >= r.MaxCapacity {
		return ErrRoomMaxCapacity
	}

	if slices.Contains(r.UserIDs, userID) {
		return ErrUserAlreadyInRoom
	}
	r.UserIDs = append(r.UserIDs, userID)
	return nil
}

func (r *Room) Leave(userID uuid.UUID) error {
	if !slices.Contains(r.UserIDs, userID) {
		return ErrUserNotInRoom
	}
	r.UserIDs = vector.New(r.UserIDs).Filter(func(u uuid.UUID) bool {
		return u != userID
	})
	return nil
}
