package private_kvs

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/usecase/private_kvs_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/google/uuid"
)

type handler struct {
	uc *private_kvs_usecase.Usecase
}

func NewHandler(uc *private_kvs_usecase.Usecase, advice advice.Advice) apiconnect.PrivateKVSServiceHandler {
	return api.NewPrivateKVSServiceHandler(&handler{uc: uc}, advice)
}

func (h handler) GetETagV1(ctx context.Context, req *advice.Request[*api.PrivateKVSServiceGetETagV1Request]) (*api.PrivateKVSServiceGetETagV1Response, error) {
	user, _ := req.Principal()
	result, err := h.uc.Get(ctx, req.TransactionContext(), user, model.NewKVSCriteria(nil, nil))
	if err != nil {
		return nil, err
	}
	if result.ETag() == uuid.Nil {
		return &api.PrivateKVSServiceGetETagV1Response{
			Etag: "",
		}, nil
	}
	return &api.PrivateKVSServiceGetETagV1Response{
		Etag: result.ETag().String(),
	}, nil
}

func (h handler) GetV1(ctx context.Context, req *advice.Request[*api.PrivateKVSServiceGetV1Request]) (*api.PrivateKVSServiceGetV1Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) SetV1(ctx context.Context, req *advice.Request[*api.PrivateKVSServiceSetV1Request]) (*api.PrivateKVSServiceSetV1Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) SetV1Errors(errs *api.PrivateKVSServiceSetV1Errors) {
}
