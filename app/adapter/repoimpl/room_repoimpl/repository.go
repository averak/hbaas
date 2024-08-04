package room_repoimpl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/infrastructure/trace"
	"github.com/averak/hbaas/pkg/vector"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct{}

func NewRepository() repository.RoomRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, roomID uuid.UUID) (model.Room, error) {
	ctx, span := trace.StartSpan(ctx, "room_repoimpl.Get")
	defer span.End()

	dto, err := dao.FindRoom(ctx, tx, roomID.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Room{}, repository.ErrRoomNotFound
		}
		return model.Room{}, err
	}

	users, err := dao.RoomUsers(dao.RoomUserWhere.RoomID.EQ(roomID.String())).All(ctx, tx)
	if err != nil {
		return model.Room{}, err
	}
	userIDs := make([]uuid.UUID, 0, len(users))
	for _, user := range users {
		userIDs = append(userIDs, uuid.MustParse(user.UserID))
	}

	return model.NewRoom(
		uuid.MustParse(dto.ID),
		uuid.MustParse(dto.OwnerUserID),
		model.RoomType(dto.Type),
		dto.MaxCapacity,
		dto.Secret,
		dto.Details,
		userIDs,
	), nil
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, room model.Room) error {
	ctx, span := trace.StartSpan(ctx, "room_repoimpl.Save")
	defer span.End()

	dto := dao.Room{
		ID:          room.ID.String(),
		OwnerUserID: room.OwnerUserID.String(),
		Type:        int(room.Type),
		MaxCapacity: room.MaxCapacity,
		Secret:      room.Secret,
		Details:     room.Details,
	}
	err := dto.Upsert(ctx, tx, true, dao.RoomPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	_, err = dao.RoomUsers(dao.RoomUserWhere.RoomID.EQ(room.ID.String())).DeleteAll(ctx, tx)
	if err != nil {
		return err
	}

	users := vector.Map(room.UserIDs, func(userID uuid.UUID) *dao.RoomUser {
		return &dao.RoomUser{
			RoomID: room.ID.String(),
			UserID: userID.String(),
		}
	})
	_, err = dao.RoomUserSlice(users).UpsertAll(ctx, tx, true, dao.RoomUserPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
