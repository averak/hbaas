package system_builder

import (
	"time"

	"github.com/averak/hbaas/app/domain/model"
)

type LeaderBoardBuilder struct {
	data model.LeaderBoard
}

func NewLeaderBoardBuilder(id string) *LeaderBoardBuilder {
	return &LeaderBoardBuilder{
		data: model.NewLeaderBoard(id, nil),
	}
}

func (b LeaderBoardBuilder) Build() model.LeaderBoard {
	return b.data
}

func (b *LeaderBoardBuilder) Scores(v ...model.LeaderBoardScore) *LeaderBoardBuilder {
	b.data.Scores = append(b.data.Scores, v...)
	return b
}

type LeaderBoardScoreBuilder struct {
	data model.LeaderBoardScore
}

func NewLeaderBoardScoreBuilder(scoreID string) *LeaderBoardScoreBuilder {
	return &LeaderBoardScoreBuilder{
		data: model.NewLeaderBoardScore(scoreID, 0, time.Now()),
	}
}

func (b LeaderBoardScoreBuilder) Build() model.LeaderBoardScore {
	return b.data
}

func (b *LeaderBoardScoreBuilder) Score(v int) *LeaderBoardScoreBuilder {
	b.data.Score = v
	return b
}

func (b *LeaderBoardScoreBuilder) Timestamp(v time.Time) *LeaderBoardScoreBuilder {
	b.data.Timestamp = v
	return b
}
