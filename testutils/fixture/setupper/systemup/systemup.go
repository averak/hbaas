package systemup

import (
	"context"
	"testing"

	"github.com/averak/hbaas/app/adapter/repoimpl/global_kvs_repoimpl"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/fixture/builder/system_builder"
)

var repos = struct {
	GlobalKVSRepo repository.GlobalKVSRepository
}{
	GlobalKVSRepo: global_kvs_repoimpl.NewRepository(),
}

func Setup(t *testing.T, ctx context.Context, data ...system_builder.Data) {
	t.Helper()

	if len(data) == 0 {
		return
	}

	conn := testutils.MustDBConn(t)
	err := conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		for _, d := range data {
			if d.GlobalKVSBucket != nil {
				err := repos.GlobalKVSRepo.Save(ctx, tx, *d.GlobalKVSBucket)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
