package testconnect

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/infrastructure/connect/error_response"
	"github.com/averak/hbaas/app/infrastructure/connect/mdval"
	"github.com/averak/hbaas/app/infrastructure/session"
	"github.com/averak/hbaas/protobuf/api/api_errors"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
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

func WithSession(t *testing.T, userID uuid.UUID) Option {
	t.Helper()

	return func(header http.Header) {
		sess := session.NewSession(userID, time.Now(), faker.MaxTime())
		token, err := session.EncodeSessionToken(sess, []byte(config.Get().GetApiServer().GetSession().GetSecretKey()))
		if err != nil {
			t.Fatal(err)
		}
		header.Add(string(mdval.SessionTokenKey), token)
	}
}

func AssertErrorCode[T error_response.CodeType](t *testing.T, want T, err error) {
	var (
		details []*connect.ErrorDetail
		ce      *connect.Error
		ee      error_response.Error
	)
	if ok := errors.As(err, &ce); ok {
		details = ce.Details()
	} else if ok := errors.As(err, &ee); ok {
		details = ee.ConnectError().Details()
	} else {
		require.Fail(t, fmt.Sprintf("unexpected error type[%T]", err))
	}

	var got T
	if len(details) > 0 {
		var errDetail api_errors.ErrorDetail
		err = proto.Unmarshal(details[0].Bytes(), &errDetail)
		require.NoErrorf(t, err, "failed to unmarshal ErrorDetail[%s]", details[0].Bytes())
		got = T(errDetail.ErrorCode)
	}
	assert.Equalf(t, want, got, "(want, got) = (%s, %s[%s])", want, got, err)
}

func GetErrorDetail(err error) *api_errors.ErrorDetail {
	var ce *connect.Error
	if ok := errors.As(err, &ce); !ok {
		return nil
	}

	for _, detail := range ce.Details() {
		var errDetail api_errors.ErrorDetail
		if err = proto.Unmarshal(detail.Bytes(), &errDetail); err == nil {
			return &errDetail
		}
	}
	return nil
}
