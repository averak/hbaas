package model

import (
	"github.com/google/uuid"
)

// LobbyService は、WebRTC のロビー機能を提供します。
type LobbyService struct{}

func NewLobbyService() LobbyService {
	return LobbyService{}
}

func (s LobbyService) CreatePublicRoom(id uuid.UUID, ownerUserID uuid.UUID, maxCapacity int, detail []byte) Room {
	return NewRoom(id, ownerUserID, RoomTypePublic, maxCapacity, "", detail, nil)
}

func (s LobbyService) CreatePrivateRoom(id uuid.UUID, ownerUserID uuid.UUID, maxCapacity int, secret string, detail []byte) (Room, error) {
	if secret == "" {
		return Room{}, ErrRoomSecretEmpty
	}
	return NewRoom(id, ownerUserID, RoomTypePrivate, maxCapacity, secret, detail, nil), nil
}
