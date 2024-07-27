package authentication_repoimpl

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

func NewRepository() repository.AuthenticationRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, userID uuid.UUID) (model.UserAuthentication, error) {
	ctx, span := trace.StartSpan(ctx, "authentication_repoimpl.Get")
	defer span.End()

	dto, err := dao.UserAuthentications(dao.UserAuthenticationWhere.UserID.EQ(userID.String())).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserAuthentication{}, repository.ErrUserAuthenticationNotFound
		}
		return model.UserAuthentication{}, err
	}
	return toModel(dto), nil
}

func (r Repository) GetByBaasUserID(ctx context.Context, tx transaction.Transaction, baasUserID string) (model.UserAuthentication, error) {
	ctx, span := trace.StartSpan(ctx, "authentication_repoimpl.GetByBaasUserID")
	defer span.End()

	dto, err := dao.UserAuthentications(dao.UserAuthenticationWhere.BaasUserID.EQ(baasUserID)).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserAuthentication{}, repository.ErrUserAuthenticationNotFound
		}
		return model.UserAuthentication{}, err
	}
	return toModel(dto), nil
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, auth model.UserAuthentication) error {
	ctx, span := trace.StartSpan(ctx, "authentication_repoimpl.Save")
	defer span.End()

	dto := dao.UserAuthentication{
		UserID:              auth.UserID.String(),
		BaasUserID:          auth.BaasUserID,
		LastAuthenticatedAt: auth.LastAuthenticatedAt,
	}
	err := dto.Upsert(ctx, tx, true, dao.UserAuthenticationPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func toModel(dto *dao.UserAuthentication) model.UserAuthentication {
	return model.NewUserAuthentication(uuid.MustParse(dto.UserID), dto.BaasUserID, dto.LastAuthenticatedAt)
}
