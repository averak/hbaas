package global_kvs_repoimpl

import (
	"context"

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

func NewRepository() repository.GlobalKVSRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, criteria model.KVSCriteria) (model.GlobalKVSBucket, error) {
	ctx, span := trace.StartSpan(ctx, "global_kvs_repoimpl.Get")
	defer span.End()

	if criteria.IsEmpty() {
		return model.NewGlobalKVSBucket(nil), nil
	}

	var or []qm.QueryMod
	if len(criteria.ExactMatch) > 0 {
		or = append(or, qm.Or2(dao.GlobalKVSEntryWhere.Key.IN(criteria.ExactMatch)))
	}
	if len(criteria.PrefixMatch) > 0 {
		for _, pm := range criteria.PrefixMatch {
			or = append(or, qm.Or2(dao.GlobalKVSEntryWhere.Key.LIKE(pm+"%")))
		}
	}
	dtos, err := dao.GlobalKVSEntries(qm.Expr(or...), qm.OrderBy("key ASC")).All(ctx, tx)
	if err != nil {
		return model.GlobalKVSBucket{}, err
	}

	entries := make([]model.KVSEntry, len(dtos))
	for i, dto := range dtos {
		entries[i], err = model.NewKVSEntry(dto.Key, dto.Value)
		if err != nil {
			return model.GlobalKVSBucket{}, err
		}
	}
	return model.NewGlobalKVSBucket(entries), nil
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, bucket model.GlobalKVSBucket) error {
	ctx, span := trace.StartSpan(ctx, "global_kvs_repoimpl.Save")
	defer span.End()

	deleted := vector.Map(
		vector.New(bucket.Raw()).Filter(func(entry model.KVSEntry) bool {
			return entry.IsEmpty()
		}),
		toDto,
	)
	_, err := dao.GlobalKVSEntrySlice(deleted).DeleteAll(ctx, tx)
	if err != nil {
		return err
	}

	upserted := vector.Map(
		vector.New(bucket.Raw()).Filter(func(entry model.KVSEntry) bool {
			return !entry.IsEmpty()
		}),
		toDto,
	)
	_, err = dao.GlobalKVSEntrySlice(upserted).UpsertAll(ctx, tx, true, dao.GlobalKVSEntryPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func toDto(entry model.KVSEntry) *dao.GlobalKVSEntry {
	return &dao.GlobalKVSEntry{
		Key:   entry.Key,
		Value: entry.Value,
	}
}
