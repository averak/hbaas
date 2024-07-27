package session

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/adapter/repoimpl/authentication_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_repoimpl"
	"github.com/averak/hbaas/app/adapter/usecaseimpl"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/infrastructure/connect/interceptor"
	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"github.com/averak/hbaas/app/infrastructure/session"
	"github.com/averak/hbaas/app/usecase/session_usecase"
	"github.com/averak/hbaas/protobuf/api"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/averak/hbaas/testutils/fixture"
	"github.com/averak/hbaas/testutils/testconnect"
	"github.com/averak/hbaas/testutils/testgoogle_cloud"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_handler_AuthorizeV1(t *testing.T) {
	var (
		validIDToken   = "valid token"
		expiredIDToken = "expired token"
	)

	now := time.Now().Truncate(time.Second)
	baasCli := testgoogle_cloud.NewFirebaseClientStub(testgoogle_cloud.WithVerifyIDTokenFn(func(idToken string) (*google_cloud.FirebaseAuthIDToken, error) {
		if idToken == validIDToken {
			return &google_cloud.FirebaseAuthIDToken{
				UID: faker.UUIDv5("bu1").String(),
				Claims: map[string]any{
					google_cloud.FirebaseAuthEmailClaimKey: faker.Email(),
				},
			}, nil
		} else if idToken == expiredIDToken {
			return nil, google_cloud.ErrFirebaseAuthExpiredIDToken
		} else {
			return nil, google_cloud.ErrFirebaseAuthInvalidIDToken
		}
	}))

	server := httptest.NewServer(newMux(t, baasCli))
	defer server.Close()

	type given struct {
		seeds []fixture.Seed
	}
	type when struct {
		req *api.SessionServiceAuthorizeV1Request
		now time.Time
	}
	type then = func(*testing.T, *connect.Response[api.SessionServiceAuthorizeV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "ユーザが存在する状態で",
			Given: given{
				seeds: []fixture.Seed{
					&dao.User{
						ID: faker.UUIDv5("u1").String(),
					},
					&dao.UserAuthentication{
						UserID:     faker.UUIDv5("u1").String(),
						BaasUserID: faker.UUIDv5("bu1").String(),
					},
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "セッションを再発行できる",
					When: when{
						req: &api.SessionServiceAuthorizeV1Request{
							IdToken: validIDToken,
						},
						now: now,
					},
					Then: func(t *testing.T, got *connect.Response[api.SessionServiceAuthorizeV1Response], err error) {
						require.NoError(t, err)

						wantSession := session.Session{
							PrincipalID: faker.UUIDv5("u1"),
							IssuedAt:    now,
							ExpiresAt:   now.Add(time.Duration(config.Get().GetSession().GetExpirationSeconds()) * time.Second),
						}
						sess, err := session.DecodeSessionToken(got.Msg.GetSessionToken(), []byte(config.Get().GetSession().GetSecretKey()), time.Now())
						if err != nil {
							t.Fatal(err)
						}
						if diff := cmp.Diff(wantSession, sess, cmpopts.IgnoreFields(session.Session{}, "PrincipalID")); diff != "" {
							t.Errorf("(-want, +got)\n%s", diff)
						}

						want := &api.SessionServiceAuthorizeV1Response{
							UserId:    faker.UUIDv5("u1").String(),
							ExpiresAt: timestamppb.New(now.Add(time.Duration(config.Get().GetSession().GetExpirationSeconds()) * time.Second)),
						}
						if diff := cmp.Diff(want, got.Msg, cmpopts.IgnoreUnexported(api.SessionServiceAuthorizeV1Response{}, timestamppb.Timestamp{}), cmpopts.IgnoreFields(api.SessionServiceAuthorizeV1Response{}, "SessionToken")); diff != "" {
							t.Errorf("(-want, +got)\n%s", diff)
						}
					},
				},
				{
					Name: "ID トークンが有効期限切れ => エラー",
					When: when{
						req: &api.SessionServiceAuthorizeV1Request{
							IdToken: expiredIDToken,
						},
						now: now,
					},
					Then: func(t *testing.T, got *connect.Response[api.SessionServiceAuthorizeV1Response], err error) {
						// TODO: エラーコードを検証する
						require.Error(t, err)
					},
				},
				{
					Name: "不正な ID トークン => エラー",
					When: when{
						req: &api.SessionServiceAuthorizeV1Request{
							IdToken: "invalid token",
						},
						now: now,
					},
					Then: func(t *testing.T, got *connect.Response[api.SessionServiceAuthorizeV1Response], err error) {
						// TODO: エラーコードを検証する
						require.Error(t, err)
					},
				},
			},
		},
		{
			Name:  "ユーザが存在しない状態で",
			Given: given{},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "新規ユーザが作成される",
					When: when{
						req: &api.SessionServiceAuthorizeV1Request{
							IdToken: validIDToken,
						},
						now: now,
					},
					Then: func(t *testing.T, got *connect.Response[api.SessionServiceAuthorizeV1Response], err error) {
						require.NoError(t, err)

						wantSession := session.Session{
							IssuedAt:  now,
							ExpiresAt: now.Add(time.Duration(config.Get().GetSession().GetExpirationSeconds()) * time.Second),
						}
						sess, err := session.DecodeSessionToken(got.Msg.GetSessionToken(), []byte(config.Get().GetSession().GetSecretKey()), time.Now())
						if err != nil {
							t.Fatal(err)
						}
						if diff := cmp.Diff(wantSession, sess, cmpopts.IgnoreFields(session.Session{}, "PrincipalID")); diff != "" {
							t.Errorf("(-want, +got)\n%s", diff)
						}

						want := &api.SessionServiceAuthorizeV1Response{
							UserId:    sess.PrincipalID.String(),
							ExpiresAt: timestamppb.New(now.Add(time.Duration(config.Get().GetSession().GetExpirationSeconds()) * time.Second)),
						}
						if diff := cmp.Diff(want, got.Msg, cmpopts.IgnoreUnexported(api.SessionServiceAuthorizeV1Response{}, timestamppb.Timestamp{}), cmpopts.IgnoreFields(api.SessionServiceAuthorizeV1Response{}, "SessionToken")); diff != "" {
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
			fixture.SetupSeeds(t, context.Background(), given.seeds...)

			resp, err := testconnect.MethodInvoke(
				apiconnect.NewSessionServiceClient(http.DefaultClient, server.URL).AuthorizeV1,
				when.req,
				testconnect.WithAdjustedTime(when.now),
			)
			then(t, resp, err)
		})
	}
}

func newMux(t *testing.T, baasCli google_cloud.FirebaseClient) *http.ServeMux {
	conn := testutils.MustDBConn(t)
	opts := connect.WithInterceptors(interceptor.New()...)
	mux := http.NewServeMux()
	mux.Handle(apiconnect.NewSessionServiceHandler(
		NewHandler(session_usecase.NewUsecase(conn, usecaseimpl.NewFirebaseIdentityVerifier(baasCli), authentication_repoimpl.NewRepository(), user_repoimpl.NewRepository())),
		opts,
	))
	return mux
}
