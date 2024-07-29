package job

import (
	"context"
	"time"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type purgeOldEchosJob struct{}

func NewPurgeOldEchos() BatchJob {
	return purgeOldEchosJob{}
}

func (j purgeOldEchosJob) Desc() string {
	return "echos テーブルの古いレコードを削除します。"
}

func (j purgeOldEchosJob) Run(ctx context.Context, tctx transaction_context.TransactionContext, conn transaction.Connection) error {
	ttl := 24 * time.Hour
	return conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		_, err := dao.Echos(dao.EchoWhere.CreatedAt.LT(tctx.Now().Add(-ttl))).DeleteAll(ctx, tx)
		return err
	})
}
