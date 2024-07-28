package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLeaderBoard(t *testing.T) {
	type args struct {
		id     string
		scores []LeaderBoardScore
	}
	tests := []struct {
		name string
		args args
		want LeaderBoard
	}{
		{
			name: "ソートされたリーダーボードが生成される",
			args: args{
				scores: []LeaderBoardScore{
					{
						ScoreID: "1",
						Score:   1,
					},
					{
						ScoreID: "2",
						Score:   2,
					},
				},
			},
			want: LeaderBoard{
				Scores: []LeaderBoardScore{
					{
						ScoreID: "2",
						Score:   2,
					},
					{
						ScoreID: "1",
						Score:   1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewLeaderBoard(tt.args.id, tt.args.scores))
		})
	}
}

func TestLeaderBoard_SubmitScore(t *testing.T) {
	now := time.Now()

	type fields struct {
		Scores []LeaderBoardScore
	}
	type args struct {
		score LeaderBoardScore
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   LeaderBoard
	}{
		{
			name: "エンティティが存在する => スコアが更新される",
			fields: fields{
				Scores: []LeaderBoardScore{
					{
						ScoreID:   "1",
						Score:     1,
						Timestamp: now.Add(-time.Hour),
					},
				},
			},
			args: args{
				score: LeaderBoardScore{
					ScoreID:   "1",
					Score:     2,
					Timestamp: now,
				},
			},
			want: LeaderBoard{
				Scores: []LeaderBoardScore{
					{
						ScoreID:   "1",
						Score:     2,
						Timestamp: now,
					},
				},
			},
		},
		{
			name: "エンティティが存在しない => スコアが追加される",
			fields: fields{
				Scores: []LeaderBoardScore{
					{
						ScoreID:   "1",
						Score:     1,
						Timestamp: now,
					},
				},
			},
			args: args{
				score: LeaderBoardScore{
					ScoreID:   "2",
					Score:     2,
					Timestamp: now,
				},
			},
			want: LeaderBoard{
				Scores: []LeaderBoardScore{
					{
						ScoreID:   "2",
						Score:     2,
						Timestamp: now,
					},
					{
						ScoreID:   "1",
						Score:     1,
						Timestamp: now,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LeaderBoard{
				Scores: tt.fields.Scores,
			}
			l.SubmitScore(tt.args.score)
			assert.Equal(t, tt.want, l)
		})
	}
}

func TestLeaderBoard_updateRanking(t *testing.T) {
	type fields struct {
		Scores []LeaderBoardScore
	}
	tests := []struct {
		name   string
		fields fields
		want   LeaderBoard
	}{
		{
			name: "スコアの降順でソートする",
			fields: fields{
				Scores: []LeaderBoardScore{
					{
						ScoreID: "1",
						Score:   1,
					},
					{
						ScoreID: "2",
						Score:   2,
					},
				},
			},
			want: LeaderBoard{
				Scores: []LeaderBoardScore{
					{
						ScoreID: "2",
						Score:   2,
					},
					{
						ScoreID: "1",
						Score:   1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LeaderBoard{
				Scores: tt.fields.Scores,
			}
			l.updateRanking()
			assert.Equal(t, tt.want, l)
		})
	}
}

func TestLeaderBoardScore_eq(t *testing.T) {
	type fields struct {
		ScoreID string
	}
	type args struct {
		other LeaderBoardScore
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "スコアIDが一致する => true",
			fields: fields{
				ScoreID: "1",
			},
			args: args{
				other: LeaderBoardScore{
					ScoreID: "1",
				},
			},
			want: true,
		},
		{
			name: "スコアIDが一致しない => false",
			fields: fields{
				ScoreID: "1",
			},
			args: args{
				other: LeaderBoardScore{
					ScoreID: "2",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LeaderBoardScore{
				ScoreID: tt.fields.ScoreID,
			}
			assert.Equal(t, tt.want, s.eq(tt.args.other))
		})
	}
}
