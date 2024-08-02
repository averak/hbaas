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
func (l *LeaderBoard) SubmitScore(score LeaderBoardScore) {
	var exists bool
	for i := range l.Scores {
		if l.Scores[i].eq(score) {
			l.Scores[i].update(score.Score, score.Timestamp)
			exists = true
		}
	}
	if !exists {
		l.Scores = append(l.Scores, score)
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

func (s LeaderBoardScore) eq(other LeaderBoardScore) bool {
	return s.ScoreID == other.ScoreID
}
