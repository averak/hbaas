package user

import (
	"context"

	"github.com/averak/hbaas/app/adapter/pbconv"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/usecase/user_usecase"
	"github.com/averak/hbaas/pkg/vector"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/google/uuid"
)

type handler struct {
	uc *user_usecase.Usecase
}

func NewHandler(uc *user_usecase.Usecase, advice advice.Advice) apiconnect.UserServiceHandler {
	return api.NewUserServiceHandler(&handler{uc: uc}, advice)
}

func (h handler) ActivateV1(ctx context.Context, req *advice.Request[*api.UserServiceActivateV1Request]) (*api.UserServiceActivateV1Response, error) {
	user, _ := req.Principal()
	err := h.uc.Activate(ctx, user)
	if err != nil {
		return nil, err
	}
	return &api.UserServiceActivateV1Response{}, nil
}

func (h handler) ActivateV1Errors(errs *api.UserServiceActivateV1Errors) {
	errs.Map(model.ErrUserDeactivated, errs.RESOURCE_CONFLICT)
}

func (h handler) SearchProfilesV1(ctx context.Context, req *advice.Request[*api.UserServiceSearchProfilesV1Request]) (*api.UserServiceSearchProfilesV1Response, error) {
	userIDs := make([]uuid.UUID, 0, len(req.Msg().GetUserIds()))
	for _, id := range req.Msg().GetUserIds() {
		userID, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}
	result, err := h.uc.SearchProfiles(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	return &api.UserServiceSearchProfilesV1Response{
		Profiles: vector.Map(result, pbconv.ToProfilePb),
	}, nil
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
