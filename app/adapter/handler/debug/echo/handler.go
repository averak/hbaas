package echo

import (
	"context"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/core/ctxval"
	"github.com/averak/hbaas/app/usecase/echo_usecase"
	"github.com/averak/hbaas/protobuf/api/debug"
	"github.com/averak/hbaas/protobuf/api/debug/debugconnect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	uc *echo_usecase.Usecase
}

func New(uc *echo_usecase.Usecase) debugconnect.EchoServiceHandler {
	return handler{uc: uc}
}

func (h handler) EchoV1(ctx context.Context, c *connect.Request[debug.EchoServiceEchoV1Request]) (*connect.Response[debug.EchoServiceEchoV1Response], error) {
	tctx, _ := ctxval.GetTransactionContext(ctx)
	result, err := h.uc.Echo(ctx, tctx, c.Msg.GetMessage())
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&debug.EchoServiceEchoV1Response{
		Message:   result.Message,
		Timestamp: timestamppb.New(result.Timestamp),
	}), nil
}
