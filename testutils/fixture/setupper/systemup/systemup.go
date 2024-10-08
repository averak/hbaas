package systemup

import (
	"context"
	"testing"

	"github.com/averak/hbaas/app/adapter/repoimpl/global_kvs_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/leader_board_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/master_data_repoimpl"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/fixture/builder/system_builder"
)

var repos = struct {
	GlobalKVSRepo   repository.GlobalKVSRepository
	LeaderBoardRepo repository.LeaderBoardRepository
	masterDataRepo  repository.MasterDataRepository
}{
	GlobalKVSRepo:   global_kvs_repoimpl.NewRepository(),
	LeaderBoardRepo: leader_board_repoimpl.NewRepository(),
	masterDataRepo:  master_data_repoimpl.NewRepository(),
}

func Setup(t *testing.T, ctx context.Context, data ...system_builder.Data) {
	t.Helper()

	if len(data) == 0 {
		return
	}

	conn := testutils.MustDBConn(t)
	err := conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		for _, d := range data {
			for _, v := range d.MasterData {
				err := repos.masterDataRepo.Save(ctx, tx, v)
				if err != nil {
					return err
				}
			}

			if d.GlobalKVSBucket != nil {
				err := repos.GlobalKVSRepo.Save(ctx, tx, *d.GlobalKVSBucket)
				if err != nil {
					return err
				}
			}

			for _, v := range d.LeaderBoard {
				err := repos.LeaderBoardRepo.Save(ctx, tx, v)
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
