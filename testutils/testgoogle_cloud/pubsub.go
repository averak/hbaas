package testgoogle_cloud

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateTopic(t *testing.T, ctx context.Context, topicID string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cli, err := google_cloud.NewPubSubClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = cli.Close() }()

	_, err = cli.CreateTopic(ctx, topicID)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return
		}
		t.Fatal(err)
	}
}

func DeleteTopic(t *testing.T, ctx context.Context, topicID string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cli, err := google_cloud.NewPubSubClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = cli.Close() }()

	err = cli.Topic(topicID).Delete(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func CreateSubscription(t *testing.T, ctx context.Context, topicID, subscriptionID string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cli, err := google_cloud.NewPubSubClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = cli.Close() }()

	_, err = cli.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
		Topic:                     cli.Topic(topicID),
		EnableExactlyOnceDelivery: true,
	})
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return
		}
		t.Fatal(err)
	}
}

func DeleteSubscription(t *testing.T, ctx context.Context, subscriptionID string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cli, err := google_cloud.NewPubSubClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = cli.Close() }()

	err = cli.Subscription(subscriptionID).Delete(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
