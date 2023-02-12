package crontask

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Shonminh/bilibee/apps/collect/api"
)

type VideoCollectTaskSchema struct {
	VideoCollectService api.VideoCollectService
}

func (schema *VideoCollectTaskSchema) CollectVideo(ctx context.Context) error {
	err := schema.VideoCollectService.CollectVideoInfo(ctx)
	if err != nil {
		return errors.Wrap(err, "CollectVideoInfo")
	}
	return nil
}
