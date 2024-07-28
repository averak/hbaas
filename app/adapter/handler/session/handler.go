package session

import (
	"context"
	"time"

	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"github.com/averak/hbaas/app/infrastructure/session"
	"github.com/averak/hbaas/app/usecase/session_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	uc *session_usecase.Usecase
}

func NewHandler(uc *session_usecase.Usecase, advice advice.Advice) apiconnect.SessionServiceHandler {
	return api.NewSessionServiceHandler(&handler{uc: uc}, advice)
}

func (h handler) AuthorizeV1(ctx context.Context, req *advice.Request[*api.SessionServiceAuthorizeV1Request]) (*api.SessionServiceAuthorizeV1Response, error) {
	result, err := h.uc.Authorize(ctx, req.TransactionContext(), req.Msg().GetIdToken())
	if err != nil {
		return nil, err
	}

	expiresAt := result.AuthorizedAt.Add(time.Duration(config.Get().GetSession().GetExpirationSeconds()) * time.Second)
	sess := session.NewSession(result.UserID, result.AuthorizedAt, expiresAt)
	token, err := session.EncodeSessionToken(sess, []byte(config.Get().GetSession().GetSecretKey()))
	if err != nil {
		return nil, err
	}
	return &api.SessionServiceAuthorizeV1Response{
		UserId:       result.UserID.String(),
		SessionToken: token,
		ExpiresAt:    timestamppb.New(expiresAt),
	}, nil
}

func (h handler) AuthorizeV1Errors(errs *api.SessionServiceAuthorizeV1Errors) {
	errs.Map(google_cloud.ErrFirebaseAuthInvalidIDToken, errs.ILLEGAL_ARGUMENT)
	errs.Map(google_cloud.ErrFirebaseAuthExpiredIDToken, errs.ID_TOKEN_EXPIRED)
}
