package private_kvs_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/registry"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/averak/hbaas/protobuf/resource"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/averak/hbaas/testutils/fixture/builder/user_builder"
	"github.com/averak/hbaas/testutils/fixture/setupper/userup"
	"github.com/averak/hbaas/testutils/testconnect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_handler_GetETagV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	type given struct {
		userData user_builder.Data
	}
	type when struct {
		req *api.PrivateKVSServiceGetETagV1Request
	}
	type then = func(*testing.T, *connect.Response[api.PrivateKVSServiceGetETagV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "ETag が存在する状態で",
			Given: given{
				userData: user_builder.New(faker.UUIDv5("u1")).
					PrivateKVSBucket(
						user_builder.NewPrivateKVSBucketBuilder(faker.UUIDv5("u1")).
							ETag(faker.UUIDv5("e1")).
							Build(),
					).
					Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "ETag を取得できる",
					When: when{
						req: &api.PrivateKVSServiceGetETagV1Request{},
					},
					Then: func(t *testing.T, got *connect.Response[api.PrivateKVSServiceGetETagV1Response], err error) {
						require.NoError(t, err)

						want := &api.PrivateKVSServiceGetETagV1Response{
							Etag: faker.UUIDv5("e1").String(),
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
			},
		},
		{
			Name: "ETag が存在しない状態で",
			Given: given{
				userData: user_builder.New(faker.UUIDv5("u1")).
					PrivateKVSBucket(
						user_builder.NewPrivateKVSBucketBuilder(faker.UUIDv5("u1")).
							ETag(uuid.Nil).
							Build(),
					).
					Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "ETag は空文字列になる",
					When: when{
						req: &api.PrivateKVSServiceGetETagV1Request{},
					},
					Then: func(t *testing.T, got *connect.Response[api.PrivateKVSServiceGetETagV1Response], err error) {
						require.NoError(t, err)

						want := &api.PrivateKVSServiceGetETagV1Response{
							Etag: "",
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			userup.Setup(t, context.Background(), given.userData)
			defer testutils.Teardown(t)

			got, err := testconnect.MethodInvoke(
				apiconnect.NewPrivateKVSServiceClient(http.DefaultClient, server.URL).GetETagV1,
				when.req,
				testconnect.WithSession(t, faker.UUIDv5("u1")),
			)
			then(t, got, err)
		})
	}
}

func Test_handler_GetV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	type given struct {
		userData user_builder.Data
	}
	type when struct {
		req *api.PrivateKVSServiceGetV1Request
	}
	type then = func(*testing.T, *connect.Response[api.PrivateKVSServiceGetV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "KVS が初期化済みの状態で",
			Given: given{
				userData: user_builder.New(faker.UUIDv5("u1")).
					PrivateKVSBucket(
						user_builder.NewPrivateKVSBucketBuilder(faker.UUIDv5("u1")).
							ETag(faker.UUIDv5("e1")).
							Entries(
								user_builder.NewKVSEntryBuilder(t).Key("group1:key1").Value([]byte("v1")).Build(),
								user_builder.NewKVSEntryBuilder(t).Key("group1:key2").Value([]byte("v2")).Build(),
								user_builder.NewKVSEntryBuilder(t).Key("group2:key1").Value([]byte("v3")).Build(),
								user_builder.NewKVSEntryBuilder(t).Key("group2:key2").Value([]byte("v4")).Build(),
							).
							Build(),
					).
					Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "前方一致で検索できる",
					When: when{
						req: &api.PrivateKVSServiceGetV1Request{
							Criteria: []*resource.KVSCriterion{
								{
									Key:          "group1",
									MatchingType: resource.KVSCriterion_MATCHING_TYPE_PREFIX_MATCH,
								},
							},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.PrivateKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.PrivateKVSServiceGetV1Response{
							Entries: []*resource.KVSEntry{
								{
									Key:   "group1:key1",
									Value: []byte("v1"),
								},
								{
									Key:   "group1:key2",
									Value: []byte("v2"),
								},
							},
							Etag: faker.UUIDv5("e1").String(),
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "完全一致で検索できる",
					When: when{
						req: &api.PrivateKVSServiceGetV1Request{
							Criteria: []*resource.KVSCriterion{
								{
									Key:          "group1:key1",
									MatchingType: resource.KVSCriterion_MATCHING_TYPE_EXACT_MATCH,
								},
								{
									Key:          "group2:key1",
									MatchingType: resource.KVSCriterion_MATCHING_TYPE_EXACT_MATCH,
								},
							},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.PrivateKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.PrivateKVSServiceGetV1Response{
							Entries: []*resource.KVSEntry{
								{
									Key:   "group1:key1",
									Value: []byte("v1"),
								},
								{
									Key:   "group2:key1",
									Value: []byte("v3"),
								},
							},
							Etag: faker.UUIDv5("e1").String(),
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "前方一致と完全一致の OR 検索で検索できる",
					When: when{
						req: &api.PrivateKVSServiceGetV1Request{
							Criteria: []*resource.KVSCriterion{
								{
									Key:          "group1",
									MatchingType: resource.KVSCriterion_MATCHING_TYPE_PREFIX_MATCH,
								},
								{
									Key:          "group2:key1",
									MatchingType: resource.KVSCriterion_MATCHING_TYPE_EXACT_MATCH,
								},
							},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.PrivateKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.PrivateKVSServiceGetV1Response{
							Entries: []*resource.KVSEntry{
								{
									Key:   "group1:key1",
									Value: []byte("v1"),
								},
								{
									Key:   "group1:key2",
									Value: []byte("v2"),
								},
								{
									Key:   "group2:key1",
									Value: []byte("v3"),
								},
							},
							Etag: faker.UUIDv5("e1").String(),
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "検索条件が空 => 空リストを返す",
					When: when{
						req: &api.PrivateKVSServiceGetV1Request{
							Criteria: []*resource.KVSCriterion{},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.PrivateKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.PrivateKVSServiceGetV1Response{
							Entries: nil,
							Etag:    faker.UUIDv5("e1").String(),
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
			},
		},
		{
			Name: "KVS が初期化されていない状態で",
			Given: given{
				userData: user_builder.New(faker.UUIDv5("u1")).Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "ETag は空文字列になる",
					When: when{
						req: &api.PrivateKVSServiceGetV1Request{
							Criteria: []*resource.KVSCriterion{},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.PrivateKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.PrivateKVSServiceGetV1Response{
							Entries: nil,
							Etag:    "",
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			userup.Setup(t, context.Background(), given.userData)
			defer testutils.Teardown(t)

			got, err := testconnect.MethodInvoke(
				apiconnect.NewPrivateKVSServiceClient(http.DefaultClient, server.URL).GetV1,
				when.req,
				testconnect.WithSession(t, faker.UUIDv5("u1")),
			)
			then(t, got, err)
		})
	}
}
