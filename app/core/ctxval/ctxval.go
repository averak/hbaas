package ctxval

import (
	"context"

	"github.com/averak/hbaas/app/core/transaction_context"
)

type (
	actionContextKey struct{}
	traceIDKey       struct{}
)

func GetTransactionContext(ctx context.Context) (transaction_context.TransactionContext, bool) {
	v, ok := ctx.Value(actionContextKey{}).(transaction_context.TransactionContext)
	return v, ok
}

func SetTransactionContext(ctx context.Context, actx transaction_context.TransactionContext) context.Context {
	return context.WithValue(ctx, actionContextKey{}, actx)
}

func GetTraceID(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(traceIDKey{}).(string)
	return v, ok
}

func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}
