package user_repoimpl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/infrastructure/trace"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct{}

func NewRepository() repository.UserRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, userID uuid.UUID) (model.User, error) {
	ctx, span := trace.StartSpan(ctx, "user_repoimpl.Get")
	defer span.End()

	dto, err := dao.Users(dao.UserWhere.ID.EQ(userID.String())).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrUserNotFound
		}
		return model.User{}, err
	}
	return model.NewUser(uuid.MustParse(dto.ID), dto.Email, model.UserStatus(dto.Status)), nil
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, user model.User) error {
	ctx, span := trace.StartSpan(ctx, "user_repoimpl.Save")
	defer span.End()

	dto := dao.User{
		ID:        user.ID.String(),
		Email:     user.Email,
		Status:    int(user.Status),
		IsDeleted: user.Status == model.UserStatusDeactivated,
	}
	err := dto.Upsert(ctx, tx, true, dao.UserPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
