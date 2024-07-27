package global_kvs_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/core/numunit"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/registry"
	api "github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/averak/hbaas/protobuf/resource"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
	"github.com/averak/hbaas/testutils/fixture/builder/system_builder"
	"github.com/averak/hbaas/testutils/fixture/setupper/systemup"
	"github.com/averak/hbaas/testutils/testconnect"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_handler_GetV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	type given struct {
		systemData system_builder.Data
	}
	type when struct {
		req *api.GlobalKVSServiceGetV1Request
	}
	type then = func(*testing.T, *connect.Response[api.GlobalKVSServiceGetV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "データが存在する状態で",
			Given: given{
				systemData: system_builder.New().
					GlobalKVSBucket(
						system_builder.NewGlobalKVSBucketBuilder().
							Entries(
								system_builder.NewKVSEntryBuilder(t).Key("group1:key1").Value([]byte("v1")).Build(),
								system_builder.NewKVSEntryBuilder(t).Key("group1:key2").Value([]byte("v2")).Build(),
								system_builder.NewKVSEntryBuilder(t).Key("group2:key1").Value([]byte("v3")).Build(),
								system_builder.NewKVSEntryBuilder(t).Key("group2:key2").Value([]byte("v4")).Build(),
							).
							Build(),
					).
					Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "前方一致で検索できる",
					When: when{
						req: &api.GlobalKVSServiceGetV1Request{
							Criteria: []*resource.KVSCriterion{
								{
									Key:          "group1",
									MatchingType: resource.KVSCriterion_MATCHING_TYPE_PREFIX_MATCH,
								},
							},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.GlobalKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.GlobalKVSServiceGetV1Response{
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
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "完全一致で検索できる",
					When: when{
						req: &api.GlobalKVSServiceGetV1Request{
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
					Then: func(t *testing.T, got *connect.Response[api.GlobalKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.GlobalKVSServiceGetV1Response{
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
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "前方一致と完全一致の OR 検索で検索できる",
					When: when{
						req: &api.GlobalKVSServiceGetV1Request{
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
					Then: func(t *testing.T, got *connect.Response[api.GlobalKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.GlobalKVSServiceGetV1Response{
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
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "検索条件が空 => 空リストを返す",
					When: when{
						req: &api.GlobalKVSServiceGetV1Request{
							Criteria: []*resource.KVSCriterion{},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.GlobalKVSServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.GlobalKVSServiceGetV1Response{
							Entries: nil,
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
				apiconnect.NewGlobalKVSServiceClient(http.DefaultClient, server.URL).GetV1,
				when.req,
			)
			then(t, got, err)
		})
	}
}

func Test_handler_SetV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	conn := testutils.MustDBConn(t)

	type given struct {
		systemData system_builder.Data
	}
	type when struct {
		req *api.GlobalKVSServiceSetV1Request
	}
	type then = func(*testing.T, *connect.Response[api.GlobalKVSServiceSetV1Response], []*dao.GlobalKVSEntry, error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "データが存在する状態で",
			Given: given{
				systemData: system_builder.New().
					GlobalKVSBucket(
						system_builder.NewGlobalKVSBucketBuilder().
							Entries(
								system_builder.NewKVSEntryBuilder(t).Key("key1").Value([]byte("v1")).Build(),
								system_builder.NewKVSEntryBuilder(t).Key("key2").Value([]byte("v2")).Build(),
								system_builder.NewKVSEntryBuilder(t).Key("key3").Value([]byte("v3")).Build(),
							).
							Build(),
					).
					Build(),
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "エントリを上書きできる",
					When: when{
						req: &api.GlobalKVSServiceSetV1Request{
							Entries: []*resource.KVSEntry{
								{
									Key:   "key2",
									Value: []byte("updated v2"),
								},
								{
									Key:   "key3",
									Value: []byte{},
								},
								{
									Key:   "key4",
									Value: []byte("inserted v4"),
								},
							},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.GlobalKVSServiceSetV1Response], dtos []*dao.GlobalKVSEntry, err error) {
						require.NoError(t, err)

						want := &api.GlobalKVSServiceSetV1Response{}
						assert.EqualExportedValues(t, want, got.Msg)

						wantDtos := []*dao.GlobalKVSEntry{
							{
								Key:   "key1",
								Value: []byte("v1"),
							},
							{
								Key:   "key2",
								Value: []byte("updated v2"),
							},
							{
								Key:   "key4",
								Value: []byte("inserted v4"),
							},
						}
						if diff := cmp.Diff(wantDtos, dtos, cmpopts.IgnoreFields(dao.GlobalKVSEntry{}, "CreatedAt", "UpdatedAt")); diff != "" {
							t.Errorf("(-want, +got)\n%s", diff)
						}
					},
				},
				{
					Name: "エントリが空 => 何もしない",
					When: when{
						req: &api.GlobalKVSServiceSetV1Request{
							Entries: []*resource.KVSEntry{},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.GlobalKVSServiceSetV1Response], dtos []*dao.GlobalKVSEntry, err error) {
						require.NoError(t, err)

						want := &api.GlobalKVSServiceSetV1Response{}
						assert.EqualExportedValues(t, want, got.Msg)

						wantDtos := []*dao.GlobalKVSEntry{
							{
								Key:   "key1",
								Value: []byte("v1"),
							},
							{
								Key:   "key2",
								Value: []byte("v2"),
							},
							{
								Key:   "key3",
								Value: []byte("v3"),
							},
						}
						if diff := cmp.Diff(wantDtos, dtos, cmpopts.IgnoreFields(dao.GlobalKVSEntry{}, "CreatedAt", "UpdatedAt")); diff != "" {
							t.Errorf("(-want, +got)\n%s", diff)
						}
					},
				},
				{
					Name: "バイナリサイズ > 100KiB の場合 => エラー",
					When: when{
						req: &api.GlobalKVSServiceSetV1Request{
							Entries: []*resource.KVSEntry{
								{
									Key:   "key1",
									Value: make([]byte, 100*numunit.KiB+1),
								},
							},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.GlobalKVSServiceSetV1Response], dtos []*dao.GlobalKVSEntry, err error) {
						// TODO: エラーコードを検証する
						require.Error(t, err)
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
				apiconnect.NewGlobalKVSServiceClient(http.DefaultClient, server.URL).SetV1,
				when.req,
			)

			var dtos []*dao.GlobalKVSEntry
			eerr := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				var eerr error
				dtos, eerr = dao.GlobalKVSEntries().All(ctx, tx)
				return eerr
			})
			if eerr != nil {
				t.Fatal(eerr)
			}

			then(t, got, dtos, err)
		})
	}
}
