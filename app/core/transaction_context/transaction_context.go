package transaction_context

import (
	"time"

	"github.com/google/uuid"
)

// TransactionContext は、機能によらずアプリケーション横断的なコンテキストを提供します。
type TransactionContext struct {
	idempotencyKey uuid.UUID
	now            time.Time
}

func NewTransactionContext(idempotencyKey uuid.UUID, now time.Time) TransactionContext {
	return TransactionContext{
		idempotencyKey: idempotencyKey,
		now:            now,
	}
}

// IdempotencyKey はトランザクションの冪等性を保証するために利用される、一意な識別子です。
// https://developer.mozilla.org/ja/docs/Glossary/Idempotent
func (c TransactionContext) IdempotencyKey() uuid.UUID {
	return c.idempotencyKey
}

func (c TransactionContext) Now() time.Time {
	return c.now
}
