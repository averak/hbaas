package main

import (
	"context"
	"errors"

	"github.com/averak/hbaas/app/core/logger"
	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"github.com/averak/hbaas/app/infrastructure/trace"
	pb "github.com/averak/hbaas/protobuf/cloud_pubsub"
)

func deleteBaasUser(ctx context.Context, firebaseCli google_cloud.FirebaseClient, msgID string, msg *pb.BaasUserDeletion) error {
	ctx, span := trace.StartSpan(ctx, "async_worker.deleteBaasUser")
	defer span.End()

	// Cloud Pub/Sub レイヤーでリトライできるので、ここではリトライしない。
	err := firebaseCli.DeleteUser(ctx, msg.GetBaasUserId())
	if err != nil {
		// ユーザが存在しない => 削除済みとみなす。
		if errors.Is(err, google_cloud.ErrFirebaseAuthUserNotFound) {
			logger.Warning(ctx, map[string]interface{}{
				"message":    "BaaS user not found",
				"messageID":  msgID,
				"baasUserId": msg.GetBaasUserId(),
			})
			return nil
		}
		return err
	}
	return nil
}
