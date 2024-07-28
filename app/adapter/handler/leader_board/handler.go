package leader_board

import (
	"context"

	"github.com/averak/hbaas/app/adapter/pbconv"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/usecase/leader_board_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
)

type handler struct {
	uc *leader_board_usecase.Usecase
}

func NewHandler(uc *leader_board_usecase.Usecase, advice advice.Advice) apiconnect.LeaderBoardServiceHandler {
	return api.NewLeaderBoardServiceHandler(&handler{uc: uc}, advice)
}

func (h handler) GetV1(ctx context.Context, req *advice.Request[*api.LeaderBoardServiceGetV1Request]) (*api.LeaderBoardServiceGetV1Response, error) {
	result, err := h.uc.Get(ctx, req.Msg().GetLeaderBoardId())
	if err != nil {
		return nil, err
	}
	return &api.LeaderBoardServiceGetV1Response{
		LeaderBoard: pbconv.ToLeaderBoardPb(result),
	}, nil
}

func (h handler) SubmitScoreV1(ctx context.Context, req *advice.Request[*api.LeaderBoardServiceSubmitScoreV1Request]) (*api.LeaderBoardServiceSubmitScoreV1Response, error) {
	//TODO implement me
	panic("implement me")
}
