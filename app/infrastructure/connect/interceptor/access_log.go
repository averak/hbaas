package interceptor

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/core/logger"
	"github.com/averak/hbaas/app/infrastructure/connect/mdval"
)

type (
	hintKey struct{}
	hint    = map[string]any
)

func NewAccessLogInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			begin := time.Now()
			ctx = context.WithValue(ctx, hintKey{}, make(hint))

			resp, err := next(ctx, req)

			payload := map[string]interface{}{
				"procedure":     req.Spec().Procedure,
				"requestedAt":   time.Now().UTC().Format(time.RFC3339Nano),
				"elapsedTimeMs": time.Since(begin).Milliseconds(),
				"request":       req.Any(),
				"response":      resp.Any(),
				"incomingMD":    mdval.NewIncomingMD(req.Header()),
			}
			hnt, ok := getLogHint(ctx)
			if ok {
				payload["hint"] = hnt
			}

			if err == nil {
				logger.Info(ctx, payload)
			} else {
				payload["error"] = err.Error()

				if errors.Is(ctx.Err(), context.Canceled) {
					// クライアントが切断した場合は Warning ログを出す。
					logger.Warning(ctx, payload)
				} else {
					logger.Error(ctx, payload)
				}
			}
			return resp, err
		}
	}
}

func getLogHint(ctx context.Context) (hint, bool) {
	v, ok := ctx.Value(hintKey{}).(hint)
	return v, ok
}

func AddLogHint(ctx context.Context, key string, value any) {
	if h, ok := getLogHint(ctx); ok {
		h[key] = value
	}
}
