package user_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/registry"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
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
	"github.com/stretchr/testify/require"
)

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
