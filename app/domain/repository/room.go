package repository

import (
	"context"
	"errors"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/google/uuid"
)

var (
	ErrRoomNotFound = errors.New("room not found")
)

type RoomRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, roomID uuid.UUID) (model.Room, error)
	Save(ctx context.Context, tx transaction.Transaction, room model.Room) error
}
