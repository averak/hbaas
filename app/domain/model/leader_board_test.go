package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLeaderBoard(t *testing.T) {
	type args struct {
		eventID  string
		entities []LeaderBoardEntity
	}
	tests := []struct {
		name string
		args args
		want LeaderBoard
	}{
		{
			name: "ソートされたリーダーボードが生成される",
			args: args{
				entities: []LeaderBoardEntity{
					{
						EntityID: "1",
						Score:    1,
					},
					{
						EntityID: "2",
						Score:    2,
					},
				},
			},
			want: LeaderBoard{
				Entries: []LeaderBoardEntity{
					{
						EntityID: "2",
						Score:    2,
					},
					{
						EntityID: "1",
						Score:    1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewLeaderBoard(tt.args.eventID, tt.args.entities), "NewLeaderBoard(%v, %v)", tt.args.eventID, tt.args.entities)
		})
	}
}

func TestLeaderBoard_SubmitScore(t *testing.T) {
	now := time.Now()

	type fields struct {
		Entries []LeaderBoardEntity
	}
	type args struct {
		entityID string
		score    int
		now      time.Time
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
				Entries: []LeaderBoardEntity{
					{
						EntityID:  "1",
						Score:     1,
						Timestamp: now.Add(-time.Hour),
					},
				},
			},
			args: args{
				entityID: "1",
				score:    2,
				now:      now,
			},
			want: LeaderBoard{
				Entries: []LeaderBoardEntity{
					{
						EntityID:  "1",
						Score:     2,
						Timestamp: now,
					},
				},
			},
		},
		{
			name: "エンティティが存在しない => スコアが追加される",
			fields: fields{
				Entries: []LeaderBoardEntity{
					{
						EntityID:  "1",
						Score:     1,
						Timestamp: now,
					},
				},
			},
			args: args{
				entityID: "2",
				score:    2,
				now:      now,
			},
			want: LeaderBoard{
				Entries: []LeaderBoardEntity{
					{
						EntityID:  "2",
						Score:     2,
						Timestamp: now,
					},
					{
						EntityID:  "1",
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
				Entries: tt.fields.Entries,
			}
			l.SubmitScore(tt.args.entityID, tt.args.score, tt.args.now)
			assert.Equal(t, tt.want, l)
		})
	}
}

func TestLeaderBoard_updateRanking(t *testing.T) {
	type fields struct {
		EventID string
		Entries []LeaderBoardEntity
	}
	tests := []struct {
		name   string
		fields fields
		want   LeaderBoard
	}{
		{
			name: "スコアの降順でソートする",
			fields: fields{
				Entries: []LeaderBoardEntity{
					{
						EntityID: "1",
						Score:    1,
					},
					{
						EntityID: "2",
						Score:    2,
					},
				},
			},
			want: LeaderBoard{
				Entries: []LeaderBoardEntity{
					{
						EntityID: "2",
						Score:    2,
					},
					{
						EntityID: "1",
						Score:    1,
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LeaderBoard{
				EventID: tt.fields.EventID,
				Entries: tt.fields.Entries,
			}
			l.updateRanking()
			assert.Equal(t, tt.want, l)
		})
	}
}
