package pbconv

import (
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/pkg/vector"
	"github.com/averak/hbaas/protobuf/resource"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToLeaderBoardPb(m model.LeaderBoard) *resource.LeaderBoard {
	return &resource.LeaderBoard{
		LeaderBoardId: m.ID,
		Scores:        vector.Map(m.Scores, toLeaderBoardScorePb),
	}
}

func toLeaderBoardScorePb(m model.LeaderBoardScore) *resource.LeaderBoardScore {
	return &resource.LeaderBoardScore{
		ScoreId:   m.ScoreID,
		Score:     int64(m.Score),
		Timestamp: timestamppb.New(m.Timestamp),
	}
}
