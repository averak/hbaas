package user

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/usecase/user_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
)

type handler struct {
	uc *user_usecase.Usecase
}

func NewHandler(uc *user_usecase.Usecase, advice advice.Advice) apiconnect.UserServiceHandler {
	return api.NewUserServiceHandler(&handler{uc: uc}, advice)
}

func (h handler) ActivateV1(ctx context.Context, req *advice.Request[*api.UserServiceActivateV1Request]) (*api.UserServiceActivateV1Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) ActivateV1Errors(errs *api.UserServiceActivateV1Errors) {
}

func (h handler) EditProfileV1(ctx context.Context, req *advice.Request[*api.UserServiceEditProfileV1Request]) (*api.UserServiceEditProfileV1Response, error) {
	user, _ := req.Principal()
	err := h.uc.EditProfile(ctx, user, req.Msg().GetData())
	if err != nil {
		return nil, err
	}
	return &api.UserServiceEditProfileV1Response{}, nil
}

func (h handler) EditProfileV1Errors(errs *api.UserServiceEditProfileV1Errors) {
	errs.Map(model.ErrUserProfileTooLarge, errs.ILLEGAL_ARGUMENT)
}

func (h handler) AccountDeleteV1(ctx context.Context, req *advice.Request[*api.UserServiceAccountDeleteV1Request]) (*api.UserServiceAccountDeleteV1Response, error) {
	user, _ := req.Principal()
	err := h.uc.AccountDelete(ctx, user)
	if err != nil {
		return nil, err
	}
	return &api.UserServiceAccountDeleteV1Response{}, nil
}
