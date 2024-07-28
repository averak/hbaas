package advice

import (
	"context"
	"fmt"
	"net/http"
	"net/textproto"
	"testing"
	"time"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_repoimpl"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/infrastructure/connect/mdval"
	"github.com/averak/hbaas/app/infrastructure/session"
	"github.com/averak/hbaas/protobuf/api/api_errors"
	pb "github.com/averak/hbaas/protobuf/config"
	"github.com/averak/hbaas/testutils"
	"github.com/averak/hbaas/testutils/bdd"
	"github.com/averak/hbaas/testutils/faker"
	"github.com/averak/hbaas/testutils/fixture"
	"github.com/averak/hbaas/testutils/testconnect"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_checkSession(t *testing.T) {
	now := time.Now()

	sessionToken := func(userID uuid.UUID) string {
		sessionToken, err := session.EncodeSessionToken(
			session.NewSession(userID, now, now.Add(1*time.Hour)),
			[]byte(config.Get().GetSession().GetSecretKey()),
		)
		if err != nil {
			t.Fatal(err)
		}
		return sessionToken
	}

	type given struct {
		conf  *config.Config
		seeds []fixture.Seed
	}
	type when struct {
		incomingMD mdval.IncomingMD
		now        time.Time
	}
	type then = func(*testing.T, *model.User, error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "デバッグモードの場合",
			Given: given{
				conf: &config.Config{
					Debug: true,
				},
				seeds: []fixture.Seed{
					&dao.User{
						ID:     faker.UUIDv5("u1").String(),
						Email:  "u1@example.com",
						Status: int(model.UserStatusActive),
					},
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "デバッグ用ヘッダが指定されている && ユーザが存在する場合 => ユーザ情報を取得できる",
					When: when{
						incomingMD: mdval.NewIncomingMD(http.Header{
							textproto.CanonicalMIMEHeaderKey(string(mdval.DebugSpoofingUserIDKey)): {faker.UUIDv5("u1").String()},
						}),
						now: now,
					},
					Then: func(t *testing.T, got *model.User, err error) {
						require.NoError(t, err)

						want := &model.User{
							ID:     faker.UUIDv5("u1"),
							Email:  "u1@example.com",
							Status: model.UserStatusActive,
						}
						assert.Equal(t, want, got)
					},
				},
				{
					Name: "デバッグ用ヘッダが指定されている && ユーザが存在しない場合 => ユーザを新規作成できる",
					When: when{
						incomingMD: mdval.NewIncomingMD(http.Header{
							textproto.CanonicalMIMEHeaderKey(string(mdval.DebugSpoofingUserIDKey)): {faker.UUIDv5("u2").String()},
						}),
						now: now,
					},
					Then: func(t *testing.T, got *model.User, err error) {
						require.NoError(t, err)

						want := &model.User{
							ID:     faker.UUIDv5("u2"),
							Email:  fmt.Sprintf("%s@example.com", faker.UUIDv5("u2")),
							Status: model.UserStatusActive,
						}
						assert.Equal(t, want, got)
					},
				},
			},
		},
		{
			Name: "デバッグモードではない場合",
			Given: given{
				conf: &config.Config{
					Debug: false,
					Session: &pb.Session{
						SecretKey: config.Get().GetSession().GetSecretKey(),
					},
				},
				seeds: []fixture.Seed{
					&dao.User{
						ID:     faker.UUIDv5("u1").String(),
						Email:  "u1@example.com",
						Status: int(model.UserStatusActive),
					},
					&dao.User{
						ID:     faker.UUIDv5("u2").String(),
						Email:  "u2@example.com",
						Status: int(model.UserStatusDeactivated),
					},
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "セッショントークンが指定されている && ユーザが存在する場合 => ユーザ情報を取得できる",
					When: when{
						incomingMD: mdval.NewIncomingMD(http.Header{
							textproto.CanonicalMIMEHeaderKey(string(mdval.SessionTokenKey)): {sessionToken(faker.UUIDv5("u1"))},
						}),
						now: now,
					},
					Then: func(t *testing.T, got *model.User, err error) {
						require.NoError(t, err)

						want := &model.User{
							ID:     faker.UUIDv5("u1"),
							Email:  "u1@example.com",
							Status: model.UserStatusActive,
						}
						assert.Equal(t, want, got)
					},
				},
				{
					Name: "セッショントークンが指定されている && ユーザが退会済み => エラー",
					When: when{
						incomingMD: mdval.NewIncomingMD(http.Header{
							textproto.CanonicalMIMEHeaderKey(string(mdval.SessionTokenKey)): {sessionToken(faker.UUIDv5("u2"))},
						}),
						now: now,
					},
					Then: func(t *testing.T, got *model.User, err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_COMMON_INVALID_USER_AVAILABILITY, err)
					},
				},
				{
					Name: "セッショントークンが指定されている && ユーザが存在しない場合 => エラー",
					When: when{
						incomingMD: mdval.NewIncomingMD(http.Header{
							textproto.CanonicalMIMEHeaderKey(string(mdval.SessionTokenKey)): {sessionToken(faker.UUIDv5("not_exists"))},
						}),
						now: now,
					},
					Then: func(t *testing.T, got *model.User, err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_COMMON_INVALID_SESSION, err)
					},
				},
				{
					Name: "セッショントークンが有効期限切れ => エラー",
					When: when{
						incomingMD: mdval.NewIncomingMD(http.Header{
							textproto.CanonicalMIMEHeaderKey(string(mdval.SessionTokenKey)): {sessionToken(faker.UUIDv5("u1"))},
						}),
						now: now.Add(1 * time.Hour),
					},
					Then: func(t *testing.T, got *model.User, err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_COMMON_INVALID_SESSION, err)
					},
				},
				{
					Name: "不正なセッショントークン => エラー",
					When: when{
						incomingMD: mdval.NewIncomingMD(http.Header{
							textproto.CanonicalMIMEHeaderKey(string(mdval.SessionTokenKey)): {"invalid"},
						}),
						now: now,
					},
					Then: func(t *testing.T, got *model.User, err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_COMMON_INVALID_SESSION, err)
					},
				},
				{
					Name: "デバッグ用ヘッダは無視される",
					When: when{
						incomingMD: mdval.NewIncomingMD(http.Header{
							textproto.CanonicalMIMEHeaderKey(string(mdval.DebugSpoofingUserIDKey)): {faker.UUIDv5("u1").String()},
						}),
						now: now,
					},
					Then: func(t *testing.T, got *model.User, err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_COMMON_INVALID_SESSION, err)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			fixture.SetupSeeds(t, context.Background(), given.seeds...)
			defer testutils.Teardown(t)

			conn := testutils.MustDBConn(t)

			got, err := checkSession(context.Background(), given.conf, user_repoimpl.NewRepository(), conn, when.incomingMD, when.now)
			then(t, got, err)
		})
	}
}

func Test_setupSpoofingUser(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type args struct {
		userID uuid.UUID
	}
	tests := []struct {
		name     string
		seeds    []fixture.Seed
		args     args
		want     model.User
		wantDtos []*dao.User
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "ユーザが存在する場合 => ユーザ情報を取得できる",
			seeds: []fixture.Seed{
				&dao.User{
					ID: faker.UUIDv5("u1").String(),
				},
			},
			args: args{
				userID: faker.UUIDv5("u1"),
			},
			want: model.User{
				ID: faker.UUIDv5("u1"),
			},
			wantDtos: []*dao.User{
				{
					ID: faker.UUIDv5("u1").String(),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name:  "ユーザが存在しない場合 => ユーザを作成できる",
			seeds: []fixture.Seed{},
			args: args{
				userID: faker.UUIDv5("u1"),
			},
			want: model.User{
				ID:     faker.UUIDv5("u1"),
				Email:  fmt.Sprintf("%s@example.com", faker.UUIDv5("u1")),
				Status: model.UserStatusActive,
			},
			wantDtos: []*dao.User{
				{
					ID:     faker.UUIDv5("u1").String(),
					Email:  fmt.Sprintf("%s@example.com", faker.UUIDv5("u1")),
					Status: int(model.UserStatusActive),
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture.SetupSeeds(t, context.Background(), tt.seeds...)
			defer testutils.Teardown(t)

			got, err := setupSpoofingUser(context.Background(), conn, user_repoimpl.NewRepository(), tt.args.userID)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)

			var dtos []*dao.User
			eerr := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				var err error
				dtos, err = dao.Users().All(ctx, tx)
				return err
			})
			if eerr != nil {
				t.Fatal(eerr)
			}
			if diff := cmp.Diff(tt.wantDtos, dtos, cmpopts.IgnoreFields(dao.User{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
