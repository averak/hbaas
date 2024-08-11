package user_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/core/numunit"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/registry"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/api_errors"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/averak/hbaas/protobuf/resource"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/averak/hbaas/testutils/fixture/builder/user_builder"
	"github.com/averak/hbaas/testutils/fixture/setupper/userup"
	"github.com/averak/hbaas/testutils/testconnect"
	"github.com/averak/hbaas/testutils/testgoogle_cloud"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_handler_ActivateV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	type given struct {
		userData []user_builder.Data
	}
	type when struct {
		req    *api.UserServiceActivateV1Request
		userID uuid.UUID
	}
	type then = func(*testing.T, *connect.Response[api.UserServiceActivateV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "ユーザが存在する状態で",
			Given: given{
				userData: []user_builder.Data{
					user_builder.New(faker.UUIDv5("USER_ACTIVE")).
						Status(model.UserStatusActive).
						Build(),
					user_builder.New(faker.UUIDv5("USER_PENDING")).
						Status(model.UserStatusPending).
						Build(),
					user_builder.New(faker.UUIDv5("USER_DEACTIVATED")).
						Status(model.UserStatusDeactivated).
						Build(),
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "ステータスが Active の場合 => 冪等に実行できる",
					When: when{
						req:    &api.UserServiceActivateV1Request{},
						userID: faker.UUIDv5("USER_ACTIVE"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceActivateV1Response], err error) {
						require.NoError(t, err)

						want := &api.UserServiceActivateV1Response{}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "ステータスが Pending の場合 => ステータスが Active に変更される",
					When: when{
						req:    &api.UserServiceActivateV1Request{},
						userID: faker.UUIDv5("USER_PENDING"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceActivateV1Response], err error) {
						require.NoError(t, err)

						want := &api.UserServiceActivateV1Response{}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "ステータスが Deactivated の場合 => エラー",
					When: when{
						req:    &api.UserServiceActivateV1Request{},
						userID: faker.UUIDv5("USER_DEACTIVATED"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceActivateV1Response], err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_METHOD_RESOURCE_CONFLICT, err)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			defer testutils.Teardown(t)
			userup.Setup(t, context.Background(), given.userData...)

			got, err := testconnect.MethodInvoke(
				apiconnect.NewUserServiceClient(http.DefaultClient, server.URL).ActivateV1,
				when.req,
				testconnect.WithSession(t, when.userID),
			)
			then(t, got, err)
		})
	}
}

func Test_handler_SearchProfilesV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	type given struct {
		userData []user_builder.Data
	}
	type when struct {
		req    *api.UserServiceSearchProfilesV1Request
		userID uuid.UUID
	}
	type then = func(*testing.T, *connect.Response[api.UserServiceSearchProfilesV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Given: given{
				userData: []user_builder.Data{
					user_builder.New(faker.UUIDv5("u1")).
						Status(model.UserStatusActive).
						Profile(user_builder.NewUserProfile(faker.UUIDv5("u1")).Raw([]byte("v1")).Build()).
						Build(),
					user_builder.New(faker.UUIDv5("u2")).
						Status(model.UserStatusPending).
						Profile(user_builder.NewUserProfile(faker.UUIDv5("u2")).Raw([]byte("v2")).Build()).
						Build(),
					user_builder.New(faker.UUIDv5("u3")).
						Status(model.UserStatusDeactivated).
						Profile(user_builder.NewUserProfile(faker.UUIDv5("u3")).Raw([]byte("v3")).Build()).
						Build(),
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "プロフィールを編集できる",
					When: when{
						req: &api.UserServiceSearchProfilesV1Request{
							UserIds: []string{
								faker.UUIDv5("u1").String(),
								faker.UUIDv5("u2").String(),
								faker.UUIDv5("u3").String(),
								faker.UUIDv5("not_exists").String(),
							},
						},
						userID: faker.UUIDv5("u1"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceSearchProfilesV1Response], err error) {
						require.NoError(t, err)

						want := &api.UserServiceSearchProfilesV1Response{
							Profiles: []*resource.UserProfile{
								{
									UserId: faker.UUIDv5("u1").String(),
									Data:   []byte("v1"),
								},
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
			defer testutils.Teardown(t)
			userup.Setup(t, context.Background(), given.userData...)

			got, err := testconnect.MethodInvoke(
				apiconnect.NewUserServiceClient(http.DefaultClient, server.URL).SearchProfilesV1,
				when.req,
				testconnect.WithSession(t, when.userID),
			)
			then(t, got, err)
		})
	}
}

func Test_handler_EditProfileV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	type given struct {
		userData []user_builder.Data
	}
	type when struct {
		req    *api.UserServiceEditProfileV1Request
		userID uuid.UUID
	}
	type then = func(*testing.T, *connect.Response[api.UserServiceEditProfileV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "プロフィールが存在する状態で",
			Given: given{
				userData: []user_builder.Data{
					user_builder.New(faker.UUIDv5("u1")).
						Profile(user_builder.NewUserProfile(faker.UUIDv5("u1")).Raw([]byte("v1")).Build()).
						Build(),
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "プロフィールを編集できる",
					When: when{
						req: &api.UserServiceEditProfileV1Request{
							Data: []byte("v2"),
						},
						userID: faker.UUIDv5("u1"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceEditProfileV1Response], err error) {
						require.NoError(t, err)
					},
				},
				{
					Name: "バイナリサイズ > 100 KiB の場合 => エラー",
					When: when{
						req: &api.UserServiceEditProfileV1Request{
							Data: make([]byte, 100*numunit.KiB+1),
						},
						userID: faker.UUIDv5("u1"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceEditProfileV1Response], err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_METHOD_ILLEGAL_ARGUMENT, err)
					},
				},
			},
		},
		{
			Name: "プロフィールが存在しない状態で",
			Given: given{
				userData: []user_builder.Data{
					user_builder.New(faker.UUIDv5("u1")).Build(),
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "プロフィールを編集できる",
					When: when{
						req: &api.UserServiceEditProfileV1Request{
							Data: []byte("v1"),
						},
						userID: faker.UUIDv5("u1"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceEditProfileV1Response], err error) {
						require.NoError(t, err)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			defer testutils.Teardown(t)
			userup.Setup(t, context.Background(), given.userData...)

			got, err := testconnect.MethodInvoke(
				apiconnect.NewUserServiceClient(http.DefaultClient, server.URL).EditProfileV1,
				when.req,
				testconnect.WithSession(t, when.userID),
			)
			then(t, got, err)
		})
	}
}

func Test_handler_AccountDeleteV1(t *testing.T) {
	testgoogle_cloud.CreateTopic(t, context.Background(), config.Get().GetAsyncWorker().GetPubsubTopicId())
	defer testgoogle_cloud.DeleteTopic(t, context.Background(), config.Get().GetAsyncWorker().GetPubsubTopicId())

	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	type given struct {
		userData []user_builder.Data
	}
	type when struct {
		req    *api.UserServiceAccountDeleteV1Request
		userID uuid.UUID
	}
	type then = func(*testing.T, *connect.Response[api.UserServiceAccountDeleteV1Response], *dao.User, error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "ユーザが存在する状態で",
			Given: given{
				userData: []user_builder.Data{
					user_builder.New(faker.UUIDv5("USER_ACTIVE")).
						Authentication(user_builder.NewUserAuthentication(faker.UUIDv5("USER_ACTIVE")).Build()).
						Status(model.UserStatusActive).
						Build(),
					user_builder.New(faker.UUIDv5("USER_DEACTIVATED")).
						Authentication(user_builder.NewUserAuthentication(faker.UUIDv5("USER_DEACTIVATED")).Build()).
						Status(model.UserStatusDeactivated).
						Build(),
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "アクティブなユーザを削除できる",
					When: when{
						req:    &api.UserServiceAccountDeleteV1Request{},
						userID: faker.UUIDv5("USER_ACTIVE"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceAccountDeleteV1Response], dto *dao.User, err error) {
						require.NoError(t, err)

						wantDto := &dao.User{
							ID:        faker.UUIDv5("USER_ACTIVE").String(),
							Email:     "",
							Status:    int(model.UserStatusDeactivated),
							IsDeleted: true,
						}
						if diff := cmp.Diff(wantDto, dto, cmpopts.IgnoreFields(dao.User{}, "CreatedAt", "UpdatedAt")); diff != "" {
							t.Errorf("(-want, +got)\n%s", diff)
						}
					},
				},
				{
					Name: "削除済みの場合 => 冪等に実行できる",
					When: when{
						req:    &api.UserServiceAccountDeleteV1Request{},
						userID: faker.UUIDv5("USER_DEACTIVATED"),
					},
					Then: func(t *testing.T, got *connect.Response[api.UserServiceAccountDeleteV1Response], dto *dao.User, err error) {
						require.NoError(t, err)

						wantDto := &dao.User{
							ID:        faker.UUIDv5("USER_DEACTIVATED").String(),
							Email:     "",
							Status:    int(model.UserStatusDeactivated),
							IsDeleted: true,
						}
						if diff := cmp.Diff(wantDto, dto, cmpopts.IgnoreFields(dao.User{}, "CreatedAt", "UpdatedAt")); diff != "" {
							t.Errorf("(-want, +got)\n%s", diff)
						}
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			defer testutils.Teardown(t)
			userup.Setup(t, context.Background(), given.userData...)

			got, err := testconnect.MethodInvoke(
				apiconnect.NewUserServiceClient(http.DefaultClient, server.URL).AccountDeleteV1,
				when.req,
				testconnect.WithSession(t, when.userID),
			)

			conn := testutils.MustDBConn(t)
			var dto *dao.User
			eerr := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				var eerr error
				dto, eerr = dao.Users(dao.UserWhere.ID.EQ(when.userID.String())).One(ctx, tx)
				if eerr != nil {
					return eerr
				}
				return nil
			})
			if eerr != nil {
				t.Fatal(eerr)
			}
			then(t, got, dto, err)
		})
	}
}
