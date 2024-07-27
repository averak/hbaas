package session

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/core/ctxval"
	"github.com/averak/hbaas/app/infrastructure/session"
	"github.com/averak/hbaas/app/usecase/session_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	uc *session_usecase.Usecase
}

func NewHandler(uc *session_usecase.Usecase) apiconnect.SessionServiceHandler {
	return &handler{uc: uc}
}

func (h handler) AuthorizeV1(ctx context.Context, c *connect.Request[api.SessionServiceAuthorizeV1Request]) (*connect.Response[api.SessionServiceAuthorizeV1Response], error) {
	tctx, _ := ctxval.GetTransactionContext(ctx)
	result, err := h.uc.Authorize(ctx, tctx, c.Msg.GetIdToken())
	if err != nil {
		return nil, err
	}

	expiresAt := result.AuthorizedAt.Add(time.Duration(config.Get().GetSession().GetExpirationSeconds()) * time.Second)
	sess := session.NewSession(result.UserID, result.AuthorizedAt, expiresAt)
	token, err := session.EncodeSessionToken(sess, []byte(config.Get().GetSession().GetSecretKey()))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api.SessionServiceAuthorizeV1Response{
		UserId:       result.UserID.String(),
		SessionToken: token,
		ExpiresAt:    timestamppb.New(expiresAt),
	}), nil
}
