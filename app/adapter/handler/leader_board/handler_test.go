package leader_board_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/registry"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/averak/hbaas/protobuf/resource"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/averak/hbaas/testutils/fixture/builder/system_builder"
	"github.com/averak/hbaas/testutils/fixture/setupper/systemup"
	"github.com/averak/hbaas/testutils/testconnect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_handler_GetV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	now := time.Now().Truncate(time.Millisecond)

	type given struct {
		systemData system_builder.Data
	}
	type when struct {
		req *api.LeaderBoardServiceGetV1Request
	}
	type then = func(*testing.T, *connect.Response[api.LeaderBoardServiceGetV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "リーダーボードが存在する状態で",
			Given: given{
				systemData: system_builder.New().
					LeaderBoard(
						system_builder.NewLeaderBoardBuilder(faker.UUIDv5("l1").String()).
							Scores(
								system_builder.NewLeaderBoardScoreBuilder(faker.UUIDv5("s1").String()).Score(1).Timestamp(now.Add(-1*time.Hour)).Build(),
								system_builder.NewLeaderBoardScoreBuilder(faker.UUIDv5("s2").String()).Score(2).Timestamp(now).Build(),
							).
							Build(),
					).
					Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "リーダーボードが存在する => 最新のスコアを取得できる",
					When: when{
						req: &api.LeaderBoardServiceGetV1Request{
							LeaderBoardId: faker.UUIDv5("l1").String(),
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.LeaderBoardServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.LeaderBoardServiceGetV1Response{
							LeaderBoard: &resource.LeaderBoard{
								LeaderBoardId: faker.UUIDv5("l1").String(),
								Scores: []*resource.LeaderBoardScore{
									{
										ScoreId:   faker.UUIDv5("s2").String(),
										Score:     2,
										Timestamp: timestamppb.New(now),
									},
									{
										ScoreId:   faker.UUIDv5("s1").String(),
										Score:     1,
										Timestamp: timestamppb.New(now.Add(-1 * time.Hour)),
									},
								},
							},
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "リーダーボードが存在する => 空のスコアを取得できる",
					When: when{
						req: &api.LeaderBoardServiceGetV1Request{
							LeaderBoardId: faker.UUIDv5("l2").String(),
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.LeaderBoardServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.LeaderBoardServiceGetV1Response{
							LeaderBoard: &resource.LeaderBoard{
								LeaderBoardId: faker.UUIDv5("l2").String(),
								Scores:        nil,
							},
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			systemup.Setup(t, context.Background(), given.systemData)
			defer testutils.Teardown(t)

			got, err := testconnect.MethodInvoke(
				apiconnect.NewLeaderBoardServiceClient(http.DefaultClient, server.URL).GetV1,
				when.req,
				testconnect.WithSpoofingUserID(uuid.New()),
			)
			then(t, got, err)
		})
	}
}
