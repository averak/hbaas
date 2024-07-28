package model

import (
	"time"

	"github.com/averak/hbaas/pkg/vector"
)

type LeaderBoard struct {
	EventID string
	Entries []LeaderBoardEntity
}

func NewLeaderBoard(eventID string, entities []LeaderBoardEntity) LeaderBoard {
	res := LeaderBoard{
		EventID: eventID,
		Entries: entities,
	}
	res.updateRanking()
	return res
}

// SubmitScore は、リーダーボードにスコアを提出します。
func (l *LeaderBoard) SubmitScore(entityID string, score int, now time.Time) {
	for i := range l.Entries {
		if l.Entries[i].EntityID == entityID {
			l.Entries[i].Update(score, now)
			l.updateRanking()
			return
		}
	}
	l.Entries = append(l.Entries, NewLeaderBoardEntity(entityID, score, now))
	l.updateRanking()
}

func (l *LeaderBoard) updateRanking() {
	l.Entries = vector.New(l.Entries).Sort(func(x, y LeaderBoardEntity) bool {
		return x.Score > y.Score
	})
}

// LeaderBoardEntity は、スコアが登録される任意のオブジェクトを表します。
// ユーザとは限らない点に注意してください。
type LeaderBoardEntity struct {
	EntityID  string
	Score     int
	Timestamp time.Time
}

func NewLeaderBoardEntity(entityID string, score int, timestamp time.Time) LeaderBoardEntity {
	return LeaderBoardEntity{
		EntityID:  entityID,
		Score:     score,
		Timestamp: timestamp,
	}
}

func (e *LeaderBoardEntity) Update(score int, timestamp time.Time) {
	e.Score = score
	e.Timestamp = timestamp
}
