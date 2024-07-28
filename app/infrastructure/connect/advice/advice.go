package advice

import (
	"context"
	"errors"

	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/infrastructure/connect/error_response"
	"github.com/averak/hbaas/app/infrastructure/connect/mdval"
	"github.com/averak/hbaas/protobuf/custom_option"
	"google.golang.org/protobuf/proto"
)

type (
	MethodOption                   = custom_option.MethodOption
	MethodErrDefinition            = custom_option.MethodErrorDefinition
	Method[REQ, RES proto.Message] func(context.Context, *Request[REQ]) (RES, error)

	// Advice は、RPC メソッドの実行前後に処理を挟むための関数です。
	// interceptor は MethodOption を解釈できないので、似た仕組みが別途必要になります。
	Advice func(context.Context, proto.Message, mdval.IncomingMD, *MethodInfo, func(context.Context, transaction_context.TransactionContext, *model.User, mdval.IncomingMD) (proto.Message, error)) error
)

func NewAdvice(conn transaction.Connection, userRepo repository.UserRepository) Advice {
	return func(ctx context.Context, req proto.Message, incomingMD mdval.IncomingMD, info *MethodInfo, method func(context.Context, transaction_context.TransactionContext, *model.User, mdval.IncomingMD) (proto.Message, error)) error {
		params, err := fixPreconditionParams(ctx, incomingMD)
		if err != nil {
			return err
		}
		tctx := params.TransactionContext()

		var principal *model.User
		if !info.Option().GetSkipAuthenticate() {
			principal, err = checkSession(ctx, config.Get(), userRepo, conn, incomingMD, tctx.Now())
			if err != nil {
				return err
			}
		}

		_, err = method(ctx, tctx, principal, incomingMD)
		if err != nil {
			if errDef, ok := info.FindErrorDefinition(err); ok {
				return error_response.New(errDef.GetCode(), errDef.GetSeverity(), errDef.GetMessage())
			}
			return err
		}
		return nil
	}
}

type MethodInfo struct {
	opt       *MethodOption
	errCauses map[error]*MethodErrDefinition
}

func NewMethodInfo(opt *MethodOption, errCauses map[error]*MethodErrDefinition) *MethodInfo {
	return &MethodInfo{
		opt:       opt,
		errCauses: errCauses,
	}
}

func (m *MethodInfo) Option() *MethodOption {
	return m.opt
}

func (m MethodInfo) FindErrorDefinition(err error) (*MethodErrDefinition, bool) {
	for cause, def := range m.errCauses {
		if errors.Is(err, cause) {
			return def, true
		}
	}
	return nil, false
}

type Request[T any] struct {
	msg       T
	tctx      transaction_context.TransactionContext
	principal *model.User
}

func NewRequest[T proto.Message](msg T, tctx transaction_context.TransactionContext, principal *model.User) *Request[T] {
	return &Request[T]{
		msg:       msg,
		tctx:      tctx,
		principal: principal,
	}
}

func (r Request[T]) Msg() T {
	return r.msg
}

func (r Request[T]) TransactionContext() transaction_context.TransactionContext {
	return r.tctx
}

// Principal は、認証されたユーザーを返します。
// 認証必須の API は必ず true を返すので、わざわざ戻り値をチェックする必要はありません。
func (r Request[T]) Principal() (model.User, bool) {
	if r.principal == nil {
		return model.User{}, false
	}
	return *r.principal, true
}
