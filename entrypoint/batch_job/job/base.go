package job

import (
	"context"

	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

// バッチ処理は、以下の条件を満たす必要があります。
// 1. リトライ可能であること (冪等性を保証する)。
// 2. 時間のかかる処理をレコードごとに実行しないこと (これが難しい場合は、非同期実行することを検討してください)。
type BatchJob interface {
	Desc() string
	Run(ctx context.Context, tctx transaction_context.TransactionContext, conn transaction.Connection) error
}
