package master_data_repoimpl

import (
	"context"
	"database/sql"
	"errors"
	"sort"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/infrastructure/trace"
	"github.com/averak/hbaas/pkg/vector"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Repository struct{}

func NewRepository() repository.MasterDataRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, revision int) (model.MasterData, error) {
	ctx, span := trace.StartSpan(ctx, "master_data_repoimpl.Get")
	defer span.End()

	dto, err := dao.FindMasterDatum(ctx, tx, revision)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.MasterData{}, repository.ErrMasterDataNotFound
		}
		return model.MasterData{}, err
	}
	return model.NewMasterData(dto.Revision, dto.Content, dto.IsActive, dto.Comment, dto.CreatedAt)
}

func (r Repository) GetActive(ctx context.Context, tx transaction.Transaction) (model.MasterData, error) {
	ctx, span := trace.StartSpan(ctx, "master_data_repoimpl.GetActive")
	defer span.End()

	dto, err := dao.MasterData(dao.MasterDatumWhere.IsActive.EQ(true)).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.MasterData{}, repository.ErrActiveMasterDataNotFound
		}
		return model.MasterData{}, err
	}
	return model.NewMasterData(dto.Revision, dto.Content, dto.IsActive, dto.Comment, dto.CreatedAt)
}

func (r Repository) GetRevisions(ctx context.Context, tx transaction.Transaction) ([]int, error) {
	ctx, span := trace.StartSpan(ctx, "master_data_repoimpl.GetRevisions")
	defer span.End()

	dtos, err := dao.MasterData(qm.Select(dao.MasterDatumColumns.Revision)).All(ctx, tx)
	if err != nil {
		return nil, err
	}
	res := vector.Map(dtos, func(revision *dao.MasterDatum) int {
		return revision.Revision
	})
	sort.Ints(res)
	return res, nil
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, data ...model.MasterData) error {
	ctx, span := trace.StartSpan(ctx, "master_data_repoimpl.Save")
	defer span.End()

	dtos := vector.Map(data, func(datum model.MasterData) *dao.MasterDatum {
		return &dao.MasterDatum{
			Revision:  datum.Revision,
			Content:   datum.Content,
			IsActive:  datum.IsActive,
			Comment:   datum.Comment,
			CreatedAt: datum.CreatedAt,
		}
	})
	_, err := dao.MasterDatumSlice(dtos).UpsertAll(ctx, tx, true, dao.MasterDatumPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
