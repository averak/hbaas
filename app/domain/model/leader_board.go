package model

import (
	"time"

	"github.com/averak/hbaas/pkg/vector"
)

type LeaderBoard struct {
	ID     string
	Scores []LeaderBoardScore
}

func NewLeaderBoard(id string, scored []LeaderBoardScore) LeaderBoard {
	res := LeaderBoard{
		ID:     id,
		Scores: scored,
	}
	res.updateRanking()
	return res
}

// SubmitScore は、リーダーボードにスコアを提出します。
// スコアの保持者はユーザとは限らない ので、誰でもスコアを上書きできるようになっています。
// 例えばブログのいいね数ランキングをリーダーボード機能で実現する場合、いいね送信者がスコアを提出することになります。
func (l *LeaderBoard) SubmitScore(scoreID string, score int, now time.Time) {
	var exists bool
	for i := range l.Scores {
		if l.Scores[i].ScoreID == scoreID {
			l.Scores[i].update(score, now)
			exists = true
		}
	}
	if !exists {
		l.Scores = append(l.Scores, NewLeaderBoardScore(scoreID, score, now))
	}
	l.updateRanking()
}

func (l *LeaderBoard) updateRanking() {
	l.Scores = vector.New(l.Scores).Sort(func(x, y LeaderBoardScore) bool {
		return x.Score > y.Score
	})
}

// LeaderBoardScore は、リーダーボードに登録されるスコアを表します。
type LeaderBoardScore struct {
	// ScoreID は、スコア集計対象である任意のオブジェクトIDが指定されます。
	// 例: ユーザID、ブログID
	ScoreID   string
	Score     int
	Timestamp time.Time
}

func NewLeaderBoardScore(scoreID string, score int, timestamp time.Time) LeaderBoardScore {
	return LeaderBoardScore{
		ScoreID:   scoreID,
		Score:     score,
		Timestamp: timestamp,
	}
}

func (s *LeaderBoardScore) update(score int, timestamp time.Time) {
	s.Score = score
	s.Timestamp = timestamp
}
