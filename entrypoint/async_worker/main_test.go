package main

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	pb "github.com/averak/hbaas/protobuf/cloud_pubsub"
	"github.com/averak/hbaas/testutils/testgoogle_cloud"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Test_processMessage(t *testing.T) {
	type args struct {
		m *pubsub.Message
	}
	tests := []struct {
		name    string
		args    args
		test    func(t *testing.T, deleteUserCalled bool)
		wantErr bool
	}{
		{
			name: "BAAS_USER_DELETION => BaaS ユーザを削除できる",
			args: args{
				m: &pubsub.Message{
					Data: toBytes(t, &pb.Message{
						EventType: pb.EventType_EVENT_TYPE_BAAS_USER_DELETION,
						Payload: &pb.Message_BaasUserDeletion{
							BaasUserDeletion: &pb.BaasUserDeletion{
								BaasUserId: "test",
							},
						},
					}),
				},
			},
			test: func(t *testing.T, deleteUserCalled bool) {
				require.Equal(t, deleteUserCalled, true)
			},
		},
		{
			name: "EVENT_TYPE_UNSPECIFIED => 何もしない",
			args: args{
				m: &pubsub.Message{
					Data: toBytes(t, &pb.Message{
						EventType: pb.EventType_EVENT_TYPE_UNSPECIFIED,
					}),
				},
			},
			test: func(t *testing.T, deleteUserCalled bool) {
				require.Equal(t, deleteUserCalled, false)
			},
		},
		{
			name: "不正なバイナリ => 何もしない",
			args: args{
				m: &pubsub.Message{
					Data: []byte("invalid"),
				},
			},
			test: func(t *testing.T, deleteUserCalled bool) {
				require.Equal(t, deleteUserCalled, false)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var deleteUserCalled bool
			firebaseCli := testgoogle_cloud.NewFirebaseClientStub(testgoogle_cloud.WithDeleteUserFn(
				func(ctx context.Context, uid string) error {
					deleteUserCalled = true
					return nil
				}),
			)

			fn := processMessage(firebaseCli)
			fn(context.Background(), tt.args.m)
			tt.test(t, deleteUserCalled)
		})
	}
}

func toBytes(t *testing.T, msg proto.Message) []byte {
	t.Helper()
	b, err := proto.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
