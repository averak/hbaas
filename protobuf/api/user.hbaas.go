// Code generated by github.com/averak/hbaas/cmd/protoc-gen-hbaas-server. DO NOT EDIT.
// source: api/user.proto

package api

import (
	connect "connectrpc.com/connect"
	context "context"
	connect1 "github.com/averak/hbaas/app/infrastructure/connect"
	advice "github.com/averak/hbaas/app/infrastructure/connect/advice"
	custom_option "github.com/averak/hbaas/protobuf/custom_option"
	proto "google.golang.org/protobuf/proto"
)

type hbaas_UserServiceHandler interface {
	// プロフィール設定/言語設定などの初期設定が完了したら、この API を呼び出してください。
	// ユーザがアクティベートされるまで、ユーザのプロフィールは非公開になります。
	ActivateV1(ctx context.Context, req *advice.Request[*UserServiceActivateV1Request]) (*UserServiceActivateV1Response, error)
	ActivateV1Errors(errs *UserServiceActivateV1Errors)

	// プロフィールを検索します。
	SearchProfilesV1(ctx context.Context, req *advice.Request[*UserServiceSearchProfilesV1Request]) (*UserServiceSearchProfilesV1Response, error)

	// プロフィールを編集します。
	EditProfileV1(ctx context.Context, req *advice.Request[*UserServiceEditProfileV1Request]) (*UserServiceEditProfileV1Response, error)
	EditProfileV1Errors(errs *UserServiceEditProfileV1Errors)

	// 退会処理を行います。
	// アカウントは永続的に削除され、復元することはできません。
	AccountDeleteV1(ctx context.Context, req *advice.Request[*UserServiceAccountDeleteV1Request]) (*UserServiceAccountDeleteV1Response, error)
}

type UserServiceActivateV1Errors struct {
	// The user has already been deactivated.
	RESOURCE_CONFLICT *advice.MethodErrDefinition

	causes map[error]*advice.MethodErrDefinition
}

func (e *UserServiceActivateV1Errors) Map(from error, to *advice.MethodErrDefinition) {
	e.causes[from] = to
}

type UserServiceEditProfileV1Errors struct {
	// The value bytes must be less than or equal to 1KiB.
	ILLEGAL_ARGUMENT *advice.MethodErrDefinition

	causes map[error]*advice.MethodErrDefinition
}

func (e *UserServiceEditProfileV1Errors) Map(from error, to *advice.MethodErrDefinition) {
	e.causes[from] = to
}

func NewUserServiceHandler(handler hbaas_UserServiceHandler, adv advice.Advice) hbaas_UserServiceHandlerImpl {
	service := File_api_user_proto.Services().ByName("UserService")
	causes := [4]map[error]*advice.MethodErrDefinition{{}, {}, {}, {}}
	methodOpts := [4]*advice.MethodOption{}
	for i, m := 0, service.Methods(); i < 4; i++ {
		methodOpts[i] = proto.GetExtension(m.Get(i).Options(), custom_option.E_MethodOption).(*advice.MethodOption)
	}
	handler.ActivateV1Errors(&UserServiceActivateV1Errors{
		RESOURCE_CONFLICT: methodOpts[0].GetMethodErrorDefinitions()[0],
		causes:            causes[0],
	})
	handler.EditProfileV1Errors(&UserServiceEditProfileV1Errors{
		ILLEGAL_ARGUMENT: methodOpts[2].GetMethodErrorDefinitions()[0],
		causes:           causes[2],
	})
	methodInfo := [4]*advice.MethodInfo{
		advice.NewMethodInfo(methodOpts[0], causes[0]),
		advice.NewMethodInfo(methodOpts[1], causes[1]),
		advice.NewMethodInfo(methodOpts[2], causes[2]),
		advice.NewMethodInfo(methodOpts[3], causes[3]),
	}
	return hbaas_UserServiceHandlerImpl{handler: handler, advice: adv, methodInfo: methodInfo}
}

type hbaas_UserServiceHandlerImpl struct {
	handler    hbaas_UserServiceHandler
	advice     advice.Advice
	methodInfo [4]*advice.MethodInfo
}

func (h hbaas_UserServiceHandlerImpl) ActivateV1(ctx context.Context, req *connect.Request[UserServiceActivateV1Request]) (*connect.Response[UserServiceActivateV1Response], error) {
	res, err := connect1.Execute(ctx, req.Msg, req.Header(), h.methodInfo[0], h.handler.ActivateV1, h.advice)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h hbaas_UserServiceHandlerImpl) SearchProfilesV1(ctx context.Context, req *connect.Request[UserServiceSearchProfilesV1Request]) (*connect.Response[UserServiceSearchProfilesV1Response], error) {
	res, err := connect1.Execute(ctx, req.Msg, req.Header(), h.methodInfo[1], h.handler.SearchProfilesV1, h.advice)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h hbaas_UserServiceHandlerImpl) EditProfileV1(ctx context.Context, req *connect.Request[UserServiceEditProfileV1Request]) (*connect.Response[UserServiceEditProfileV1Response], error) {
	res, err := connect1.Execute(ctx, req.Msg, req.Header(), h.methodInfo[2], h.handler.EditProfileV1, h.advice)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h hbaas_UserServiceHandlerImpl) AccountDeleteV1(ctx context.Context, req *connect.Request[UserServiceAccountDeleteV1Request]) (*connect.Response[UserServiceAccountDeleteV1Response], error) {
	res, err := connect1.Execute(ctx, req.Msg, req.Header(), h.methodInfo[3], h.handler.AccountDeleteV1, h.advice)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}
