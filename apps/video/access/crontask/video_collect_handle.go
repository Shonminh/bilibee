package crontask

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Shonminh/bilibee/apps/video/api"
)

type VideoCollectTaskSchema struct {
	VideoCollectService api.VideoInfoService
}

func (schema *VideoCollectTaskSchema) CollectVideo(ctx context.Context) error {
	err := schema.VideoCollectService.CollectVideoInfo(ctx)
	if err != nil {
		return errors.Wrap(err, "CollectVideoInfo")
	}
	return nil
}

func (schema *VideoCollectTaskSchema) SyncVideoInfoToEs(ctx context.Context) error {
	err := schema.VideoCollectService.SyncVideoInfoToEs(ctx)
	if err != nil {
		return errors.Wrap(err, "SyncVideoInfoToEs")
	}
	return nil
}
