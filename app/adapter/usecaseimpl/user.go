package usecaseimpl

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/usecase/user_usecase"
	pb "github.com/averak/hbaas/protobuf/cloud_pubsub"
	"google.golang.org/protobuf/proto"
)

type BaasUserDeletionTaskQueue struct {
	cli *pubsub.Client
}

func NewBaasUserDeletionTaskQueue(cli *pubsub.Client) user_usecase.BaasUserDeletionTaskQueue {
	return BaasUserDeletionTaskQueue{
		cli: cli,
	}
}

func (q BaasUserDeletionTaskQueue) Enqueue(ctx context.Context, auth model.UserAuthentication) error {
	msg := &pb.Message{
		EventType: pb.EventType_EVENT_TYPE_BAAS_USER_DELETION,
		Payload: &pb.Message_BaasUserDeletion{
			BaasUserDeletion: &pb.BaasUserDeletion{
				BaasUserId: auth.BaasUserID,
			},
		},
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	topic := q.cli.Topic(config.Get().GetAsyncWorker().GetPubsubTopicId())
	res := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	_, err = res.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}
