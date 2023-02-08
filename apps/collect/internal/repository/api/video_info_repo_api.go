package api

import (
	"context"

	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
)

type VideoInfoRepo interface {
	CreateVideoInfo(ctx context.Context, row model.VideoInfoTab) error
	QueryVideoInfoList(ctx context.Context, mid int64, aid *int64, limit *int) ([]model.VideoInfoTab, error)
}
