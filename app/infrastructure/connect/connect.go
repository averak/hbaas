package connect

import (
	"context"
	"net/http"

	"github.com/averak/hbaas/app/core/transaction_context"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/infrastructure/connect/mdval"
	"google.golang.org/protobuf/proto"
)

func Execute[REQ, RES proto.Message](ctx context.Context, req REQ, header http.Header, info advice.MethodInfo, method func(context.Context, *advice.Request[REQ]) (RES, error), adv advice.Advice) (RES, error) {
	var res RES
	wrap := func(ctx context.Context, tctx transaction_context.TransactionContext, principal *model.User, incomingMD mdval.IncomingMD) (proto.Message, error) {
		var err error
		res, err = method(ctx, advice.NewRequest(req, tctx, principal))
		return res, err
	}
	return res, adv(ctx, req, mdval.NewIncomingMD(header), info, wrap)
}
