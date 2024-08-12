package master_data_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/registry"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/api_errors"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/averak/hbaas/protobuf/resource"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
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
		req *api.MasterDataServiceGetV1Request
	}
	type then = func(*testing.T, *connect.Response[api.MasterDataServiceGetV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "有効なリビジョンが存在する状態で",
			Given: given{
				systemData: system_builder.New().
					MasterData(
						system_builder.NewMasterDataBuilder(1).
							Content([]byte("v1")).
							IsActive(true).
							Comment("c1").
							CreatedAt(now.Add(-1 * time.Hour)).
							Build(),
					).
					MasterData(
						system_builder.NewMasterDataBuilder(2).
							Content([]byte("v2")).
							IsActive(false).
							Comment("c2").
							CreatedAt(now).
							Build(),
					).
					Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "有効なリビジョンのマスターデータを取得できる",
					When: when{
						req: &api.MasterDataServiceGetV1Request{},
					},
					Then: func(t *testing.T, got *connect.Response[api.MasterDataServiceGetV1Response], err error) {
						require.NoError(t, err)

						// 最新のリビジョンではなく、有効なリビジョンが取得される。
						want := &api.MasterDataServiceGetV1Response{
							MasterData: &resource.MasterData{
								Revision:  1,
								Content:   []byte("v1"),
								IsActive:  true,
								Comment:   "c1",
								CreatedAt: timestamppb.New(now.Add(-1 * time.Hour)),
							},
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
			},
		},
		{
			Name: "有効なリビジョンが存在しない状態で",
			Given: given{
				systemData: system_builder.New().
					MasterData(system_builder.NewMasterDataBuilder(1).IsActive(false).Build()).
					Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "エラーを返す",
					When: when{
						req: &api.MasterDataServiceGetV1Request{},
					},
					Then: func(t *testing.T, got *connect.Response[api.MasterDataServiceGetV1Response], err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_METHOD_RESOURCE_NOT_FOUND, err)
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
				apiconnect.NewMasterDataServiceClient(http.DefaultClient, server.URL).GetV1,
				when.req,
				testconnect.WithSpoofingUserID(uuid.New()),
			)
			then(t, got, err)
		})
	}
}
