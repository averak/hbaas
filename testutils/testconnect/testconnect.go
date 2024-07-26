package testconnect

import (
	"context"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/infrastructure/connect/mdval"
	"github.com/google/uuid"
)

func MethodInvoke[REQ, RES any](
	method func(context.Context, *connect.Request[REQ]) (*connect.Response[RES], error),
	req *REQ,
	opts ...Option,
) (*connect.Response[RES], error) {
	connectReq := connect.NewRequest(req)
	for _, opt := range opts {
		opt(connectReq.Header())
	}
	return method(context.Background(), connectReq)
}

type Option = func(header http.Header)

func WithTransactionContext(tctx transaction_context.TransactionContext) Option {
	return func(header http.Header) {
		header.Add(string(mdval.IdempotencyKey), tctx.IdempotencyKey().String())
		header.Add(string(mdval.DebugAdjustedTimeKey), tctx.Now().Format(time.RFC3339Nano))
	}
}

func WithAdjustedTime(t time.Time) Option {
	return func(header http.Header) {
		header.Add(string(mdval.DebugAdjustedTimeKey), t.Format(time.RFC3339Nano))
	}
}

func WithIdempotencyKey(idempotencyKey uuid.UUID) Option {
	return func(header http.Header) {
		header.Add(string(mdval.IdempotencyKey), idempotencyKey.String())
	}
}

func WithSpoofingUserID(userID uuid.UUID) Option {
	return func(header http.Header) {
		header.Add(string(mdval.DebugSpoofingUserIDKey), userID.String())
	}
}

func WithClientVersion(version string) Option {
	return func(header http.Header) {
		header.Add(string(mdval.ClientVersionKey), version)
	}
}
