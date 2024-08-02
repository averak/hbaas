package faker

import (
	"time"

	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/google/uuid"
)

type TransactionContextBuilder struct {
	data transaction_context.TransactionContext
}

func NewTransactionContextBuilder() *TransactionContextBuilder {
	return &TransactionContextBuilder{
		data: transaction_context.NewTransactionContext(uuid.New(), time.Now()),
	}
}

func (b *TransactionContextBuilder) Build() transaction_context.TransactionContext {
	return b.data
}

func (b *TransactionContextBuilder) IdempotencyKey(idempotencyKey uuid.UUID) *TransactionContextBuilder {
	b.data = transaction_context.NewTransactionContext(idempotencyKey, b.data.Now())
	return b
}

func (b *TransactionContextBuilder) Now(now time.Time) *TransactionContextBuilder {
	b.data = transaction_context.NewTransactionContext(b.data.IdempotencyKey(), now)
	return b
}
