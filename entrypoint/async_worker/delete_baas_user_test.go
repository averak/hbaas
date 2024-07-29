package main

import (
	"context"
	"errors"
	"testing"

	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	pb "github.com/averak/hbaas/protobuf/cloud_pubsub"
	"github.com/averak/hbaas/testutils/testgoogle_cloud"
)

func Test_deleteBaasUser(t *testing.T) {
	type args struct {
		firebaseCli google_cloud.FirebaseClient
		msgID       string
		msg         *pb.BaasUserDeletion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "削除成功 => エラーなし",
			args: args{
				firebaseCli: testgoogle_cloud.NewFirebaseClientStub(
					testgoogle_cloud.WithDeleteUserFn(func(ctx context.Context, uid string) error {
						return nil
					}),
				),
				msg: &pb.BaasUserDeletion{
					BaasUserId: "bu1",
				},
			},
			wantErr: false,
		},
		{
			name: "ユーザが存在しない => エラーなし",
			args: args{
				firebaseCli: testgoogle_cloud.NewFirebaseClientStub(
					testgoogle_cloud.WithDeleteUserFn(func(ctx context.Context, uid string) error {
						return google_cloud.ErrFirebaseAuthUserNotFound
					}),
				),
				msg: &pb.BaasUserDeletion{
					BaasUserId: "bu1",
				},
			},
			wantErr: false,
		},
		{
			name: "ユーザ削除に失敗 => エラーを返す",
			args: args{
				firebaseCli: testgoogle_cloud.NewFirebaseClientStub(
					testgoogle_cloud.WithDeleteUserFn(func(ctx context.Context, uid string) error {
						return errors.New("unexpected error")
					}),
				),
				msg: &pb.BaasUserDeletion{
					BaasUserId: "bu1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			err := deleteBaasUser(ctx, tt.args.firebaseCli, tt.args.msgID, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteBaasUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
