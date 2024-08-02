package global_kvs

import (
	"context"

	"github.com/averak/hbaas/app/adapter/pbconv"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/usecase/global_kvs_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
)

type handler struct {
	uc *global_kvs_usecase.Usecase
}

func NewHandler(uc *global_kvs_usecase.Usecase, advice advice.Advice) apiconnect.GlobalKVSServiceHandler {
	return api.NewGlobalKVSServiceHandler(&handler{uc: uc}, advice)
}

func (h handler) GetV1(ctx context.Context, req *advice.Request[*api.GlobalKVSServiceGetV1Request]) (*api.GlobalKVSServiceGetV1Response, error) {
	result, err := h.uc.Get(ctx, pbconv.FromKVSCriteriaPb(req.Msg().GetCriteria()))
	if err != nil {
		return nil, err
	}
	return &api.GlobalKVSServiceGetV1Response{
		Entries: pbconv.ToKVSEntryPbs(result.Raw()),
	}, nil
}

func (h handler) SetV1(ctx context.Context, req *advice.Request[*api.GlobalKVSServiceSetV1Request]) (*api.GlobalKVSServiceSetV1Response, error) {
	entries, err := pbconv.FromKVSEntryPbs(req.Msg().GetEntries())
	if err != nil {
		return nil, err
	}
	err = h.uc.Set(ctx, entries)
	if err != nil {
		return nil, err
	}
	return &api.GlobalKVSServiceSetV1Response{}, nil
}

func (h handler) SetV1Errors(errs *api.GlobalKVSServiceSetV1Errors) {
	errs.Map(model.ErrKVSEntryValueTooLarge, errs.ILLEGAL_ARGUMENT)
}
