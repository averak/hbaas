package google_cloud

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/averak/hbaas/app/core/config"
	"google.golang.org/api/option"
)

func NewPubSubClient(ctx context.Context) (*pubsub.Client, error) {
	opts := make([]option.ClientOption, 0)
	if config.Get().GetGoogleCloud().GetPubsub().GetUseEmulator() {
		_ = os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	}
	res, err := pubsub.NewClient(ctx, config.Get().GetGoogleCloud().GetProjectId(), opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
