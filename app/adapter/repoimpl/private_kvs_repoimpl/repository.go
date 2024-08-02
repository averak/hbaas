package private_kvs_repoimpl

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
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Repository struct{}

func NewRepository() repository.PrivateKVSRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, userID uuid.UUID, criteria model.KVSCriteria) (model.PrivateKVSBucket, error) {
	ctx, span := trace.StartSpan(ctx, "private_kvs_repoimpl.Get")
	defer span.End()

	var etag uuid.UUID
	etagDto, err := dao.PrivateKVSEtags(dao.PrivateKVSEtagWhere.UserID.EQ(userID.String())).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			etag = uuid.Nil
		} else {
			return model.PrivateKVSBucket{}, err
		}
	} else {
		etag = uuid.MustParse(etagDto.Etag)
	}

	if criteria.IsEmpty() {
		return model.NewPrivateKVSBucket(userID, etag, nil), nil
	}

	var or []qm.QueryMod
	if len(criteria.ExactMatch) > 0 {
		or = append(or, qm.Or2(dao.PrivateKVSEntryWhere.Key.IN(criteria.ExactMatch)))
	}
	if len(criteria.PrefixMatch) > 0 {
		for _, pm := range criteria.PrefixMatch {
			or = append(or, qm.Or2(dao.PrivateKVSEntryWhere.Key.LIKE(pm+"%")))
		}
	}
	dtos, err := dao.PrivateKVSEntries(qm.Expr(or...), qm.OrderBy("key ASC")).All(ctx, tx)
	if err != nil {
		return model.PrivateKVSBucket{}, err
	}
	entries := make([]model.KVSEntry, len(dtos))
	for i, dto := range dtos {
		entries[i], err = model.NewKVSEntry(dto.Key, dto.Value)
		if err != nil {
			return model.PrivateKVSBucket{}, err
		}
	}
	return model.NewPrivateKVSBucket(userID, etag, entries), nil
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, bucket model.PrivateKVSBucket) error {
	ctx, span := trace.StartSpan(ctx, "private_kvs_repoimpl.Save")
	defer span.End()

	etag := &dao.PrivateKVSEtag{
		UserID: bucket.UserID.String(),
		Etag:   bucket.ETag().String(),
	}
	if bucket.ETag() == uuid.Nil {
		_, err := etag.Delete(ctx, tx)
		if err != nil {
			return err
		}
	} else {
		err := etag.Upsert(ctx, tx, true, dao.PrivateKVSEtagPrimaryKeyColumns, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}
	}

	deleted := vector.Map(
		vector.New(bucket.Raw()).Filter(func(entry model.KVSEntry) bool {
			return entry.IsEmpty()
		}),
		func(entry model.KVSEntry) *dao.PrivateKVSEntry {
			return toDto(bucket.UserID, entry)
		},
	)
	_, err := dao.PrivateKVSEntrySlice(deleted).DeleteAll(ctx, tx)
	if err != nil {
		return err
	}

	upserted := vector.Map(
		vector.New(bucket.Raw()).Filter(func(entry model.KVSEntry) bool {
			return !entry.IsEmpty()
		}),
		func(entry model.KVSEntry) *dao.PrivateKVSEntry {
			return toDto(bucket.UserID, entry)
		},
	)
	_, err = dao.PrivateKVSEntrySlice(upserted).UpsertAll(ctx, tx, true, dao.PrivateKVSEntryPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func toDto(userID uuid.UUID, entry model.KVSEntry) *dao.PrivateKVSEntry {
	return &dao.PrivateKVSEntry{
		UserID: userID.String(),
		Key:    entry.Key,
		Value:  entry.Value,
	}
}
