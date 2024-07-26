package echo_repoimpl

import (
	"context"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/infrastructure/trace"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct{}

func New() repository.EchoRepository {
	return &Repository{}
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, echos ...model.Echo) error {
	ctx, span := trace.StartSpan(ctx, "echo_repoimpl.Save")
	defer span.End()

	if len(echos) == 0 {
		return nil
	}

	dtos := make([]*dao.Echo, len(echos))
	for i, echo := range echos {
		dtos[i] = &dao.Echo{
			ID:        echo.ID.String(),
			Message:   echo.Message,
			Timestamp: echo.Timestamp,
		}
	}
	_, err := dao.EchoSlice(dtos).UpsertAll(ctx, tx, true, dao.EchoPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
