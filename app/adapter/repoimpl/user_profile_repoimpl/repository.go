package user_profile_repoimpl

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

func NewRepository() repository.UserProfileRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, userID uuid.UUID) (model.UserProfile, error) {
	ctx, span := trace.StartSpan(ctx, "user_profile_repoimpl.Get")
	defer span.End()

	dto, err := dao.FindUserProfile(ctx, tx, userID.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.NewUserProfile(userID, nil)
		}
		return model.UserProfile{}, err
	}
	return model.NewUserProfile(userID, dto.Content)
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, profile model.UserProfile) error {
	ctx, span := trace.StartSpan(ctx, "user_profile_repoimpl.Save")
	defer span.End()

	dto := dao.UserProfile{
		UserID:  profile.UserID.String(),
		Content: profile.Bytes(),
	}
	err := dto.Upsert(ctx, tx, true, dao.UserProfilePrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
