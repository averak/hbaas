package master_data

import (
	"context"

	"github.com/averak/hbaas/app/adapter/pbconv"
	"github.com/averak/hbaas/app/domain/repository"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/usecase/master_data_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
)

type handler struct {
	uc *master_data_usecase.Usecase
}

func NewHandler(uc *master_data_usecase.Usecase, advice advice.Advice) apiconnect.MasterDataServiceHandler {
	return api.NewMasterDataServiceHandler(&handler{uc: uc}, advice)
}

func (h handler) GetV1(ctx context.Context, _ *advice.Request[*api.MasterDataServiceGetV1Request]) (*api.MasterDataServiceGetV1Response, error) {
	result, err := h.uc.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &api.MasterDataServiceGetV1Response{
		MasterData: pbconv.ToMasterDataPb(result),
	}, nil
}

func (h handler) GetV1Errors(errs *api.MasterDataServiceGetV1Errors) {
	errs.Map(repository.ErrActiveMasterDataNotFound, errs.RESOURCE_NOT_FOUND)
}

func (h handler) GetRevisionV1(ctx context.Context, req *advice.Request[*api.MasterDataServiceGetRevisionV1Request]) (*api.MasterDataServiceGetRevisionV1Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetRevisionV1Errors(errs *api.MasterDataServiceGetRevisionV1Errors) {
	errs.Map(repository.ErrActiveMasterDataNotFound, errs.RESOURCE_NOT_FOUND)
}
