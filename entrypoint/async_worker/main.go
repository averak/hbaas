package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/pubsub"
	"github.com/averak/hbaas/app/core/build_info"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/core/logger"
	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"github.com/averak/hbaas/app/infrastructure/trace"
	pb "github.com/averak/hbaas/protobuf/cloud_pubsub"
	"google.golang.org/protobuf/proto"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Init("async-worker", build_info.ServerVersion())
	if config.Get().GetGoogleCloud().GetTrace().GetEnabled() {
		trace.Init(ctx, "async-worker", build_info.ServerVersion(), config.Get().GetGoogleCloud().GetTrace().GetSamplingRate())
	}

	firebaseCli, err := google_cloud.NewFirebaseClient(ctx)
	if err != nil {
		logger.Emergency(ctx, map[string]interface{}{
			"message": "Failed to create firebase client",
			"error":   err.Error(),
		})
		log.Fatal(err)
	}
	pubsubCli, err := google_cloud.NewPubSubClient(ctx)
	if err != nil {
		logger.Emergency(ctx, map[string]interface{}{
			"message": "Failed to create pubsub client",
			"error":   err.Error(),
		})
		log.Fatal(err)
	}
	defer func() { _ = pubsubCli.Close() }()

	errChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		logger.Notice(ctx, "Start async worker.")
		sub := pubsubCli.Subscription(config.Get().GetAsyncWorker().GetPubsubSubscriptionId())
		errChan <- sub.Receive(ctx, processMessage(firebaseCli))
	}()

	select {
	case err := <-errChan:
		logger.Emergency(ctx, map[string]interface{}{
			"message": "Failed to start async worker.",
			"error":   err.Error(),
		})
		cancel()
	case <-sigChan:
		logger.Notice(ctx, "Shutdown signal received, shutting down process...")
	}
}

func processMessage(firebaseCli google_cloud.FirebaseClient) func(context.Context, *pubsub.Message) {
	return func(ctx context.Context, m *pubsub.Message) {
		ctx, span := trace.StartSpan(ctx, "async_worker.processMessage")
		defer span.End()

		msg := &pb.Message{}
		if err := proto.Unmarshal(m.Data, msg); err != nil {
			logger.Error(ctx, map[string]interface{}{
				"message":   "failed to unmarshal message",
				"messageID": m.ID,
				"error":     err.Error(),
			})
			m.Nack()
			return
		}

		switch msg.GetEventType() {
		case pb.EventType_EVENT_TYPE_UNSPECIFIED:
			logger.Error(ctx, map[string]interface{}{
				"message":   "unexpected event type",
				"messageID": m.ID,
				"data":      msg,
			})
			m.Nack()
			return
		case pb.EventType_EVENT_TYPE_BAAS_USER_DELETION:
			err := deleteBaasUser(ctx, firebaseCli, m.ID, msg.GetBaasUserDeletion())
			if err != nil {
				logger.Error(ctx, map[string]interface{}{
					"message":    "failed to delete baas user",
					"messageID":  m.ID,
					"baasUserID": msg.GetBaasUserDeletion().GetBaasUserId(),
					"error":      err.Error(),
				})
				m.Nack()
				return
			}
			logger.Info(ctx, map[string]interface{}{
				"message":    "baas user deleted",
				"messageID":  m.ID,
				"baasUserID": msg.GetBaasUserDeletion().GetBaasUserId(),
			})
		}

		m.Ack()
		logger.Info(ctx, map[string]interface{}{
			"message":   "message acked",
			"messageID": m.ID,
		})
	}
}
