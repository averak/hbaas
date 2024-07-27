package interceptor

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/core/ctxval"
	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/infrastructure/connect/mdval"
	"github.com/google/uuid"
)

type PreconditionParams struct {
	RequestID      uuid.UUID
	IdempotencyKey uuid.UUID
	// リクエストを受け取った現実世界での時刻
	RequestedTime time.Time
	// サーバが現在時刻として認識する時刻 (現実世界と乖離しうる)
	Now time.Time
}

func (p PreconditionParams) TransactionContext() transaction_context.TransactionContext {
	return transaction_context.NewTransactionContext(
		p.IdempotencyKey,
		p.Now,
	)
}

func NewFixPreconditionParamsInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			incomingMD := mdval.NewIncomingMD(req.Header())
			builder := newPreconditionParamsBuilder()
			builder.requestID(uuid.New())
			fixIdempotencyKey(incomingMD, builder)
			err := fixTimestamp(config.Get(), incomingMD, builder, time.Now())
			if err != nil {
				return nil, err
			}

			params, err := builder.build()
			if err != nil {
				return nil, err
			}

			AddLogHint(ctx, "requestID", params.RequestID)
			AddLogHint(ctx, "transactionContext", params.TransactionContext())

			ctx = ctxval.SetTransactionContext(ctx, params.TransactionContext())
			return next(ctx, req)
		}
	}
}

func fixIdempotencyKey(incomingMD mdval.IncomingMD, builder *preconditionParamsBuilder) {
	idempotencyKey, ok := incomingMD.Get(mdval.IdempotencyKey)
	if ok {
		ik, err := uuid.Parse(idempotencyKey)
		if err != nil {
			builder.idempotencyKey(uuid.New())
		} else {
			builder.idempotencyKey(ik)
		}
	} else {
		builder.idempotencyKey(uuid.New())
	}
}

func fixTimestamp(conf *config.Config, incomingMD mdval.IncomingMD, builder *preconditionParamsBuilder, now time.Time) error {
	builder.requestedTime(now)
	builder.now(now)
	if conf.GetDebug() {
		adjustedTimeStr, ok := incomingMD.Get(mdval.DebugAdjustedTimeKey)
		if ok {
			adjustedTime, err := time.Parse(time.RFC3339, adjustedTimeStr)
			if err != nil {
				return err
			}
			builder.now(adjustedTime)
		}
	}
	return nil
}

// PreconditionParams の設定漏れがないか検証するために、ビルダーを定義しています。
type preconditionParamsBuilder struct {
	raw PreconditionParams
}

func newPreconditionParamsBuilder() *preconditionParamsBuilder {
	return &preconditionParamsBuilder{}
}

func (b preconditionParamsBuilder) build() (PreconditionParams, error) {
	if b.raw.RequestID == uuid.Nil {
		return PreconditionParams{}, errors.New("requestID is not set")
	}
	if b.raw.IdempotencyKey == uuid.Nil {
		return PreconditionParams{}, errors.New("idempotencyKey is not set")
	}
	if b.raw.RequestedTime.IsZero() {
		return PreconditionParams{}, errors.New("requestedTime is not set")
	}
	if b.raw.Now.IsZero() {
		return PreconditionParams{}, errors.New("now is not set")
	}
	return b.raw, nil
}

func (b *preconditionParamsBuilder) requestID(id uuid.UUID) *preconditionParamsBuilder {
	b.raw.RequestID = id
	return b
}

func (b *preconditionParamsBuilder) idempotencyKey(key uuid.UUID) *preconditionParamsBuilder {
	b.raw.IdempotencyKey = key
	return b
}

func (b *preconditionParamsBuilder) requestedTime(t time.Time) *preconditionParamsBuilder {
	b.raw.RequestedTime = t
	return b
}

func (b *preconditionParamsBuilder) now(t time.Time) *preconditionParamsBuilder {
	b.raw.Now = t
	return b
}