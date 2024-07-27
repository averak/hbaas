package global_kvs

import (
	"context"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/adapter/pbconv"
	"github.com/averak/hbaas/app/core/ctxval"
	"github.com/averak/hbaas/app/usecase/global_kvs_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
)

type handler struct {
	uc *global_kvs_usecase.Usecase
}

func NewHandler(uc *global_kvs_usecase.Usecase) apiconnect.GlobalKVSServiceHandler {
	return &handler{uc: uc}
}

func (h handler) GetV1(ctx context.Context, c *connect.Request[api.GlobalKVSServiceGetV1Request]) (*connect.Response[api.GlobalKVSServiceGetV1Response], error) {
	tctx, _ := ctxval.GetTransactionContext(ctx)
	result, err := h.uc.Get(ctx, tctx, pbconv.FromKVSCriteriaPb(c.Msg.GetCriteria()))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api.GlobalKVSServiceGetV1Response{
		Entries: pbconv.ToKVSEntryPbs(result.Raw()),
	}), nil
}

func (h handler) SetV1(ctx context.Context, c *connect.Request[api.GlobalKVSServiceSetV1Request]) (*connect.Response[api.GlobalKVSServiceSetV1Response], error) {
	tctx, _ := ctxval.GetTransactionContext(ctx)
	entries, err := pbconv.FromKVSEntryPbs(c.Msg.GetEntries())
	if err != nil {
		return nil, err
	}
	err = h.uc.Set(ctx, tctx, entries)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api.GlobalKVSServiceSetV1Response{}), nil
}
