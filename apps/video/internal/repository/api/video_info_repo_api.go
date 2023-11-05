package api

import (
	"context"

	"github.com/Shonminh/bilibee/apps/video/internal/repository/model"
)

type VideoInfoRepo interface {
	BatchCreateVideoInfos(ctx context.Context, row []model.VideoInfoTab) error
	QueryVideoInfoList(ctx context.Context, mid int64, aid *int64, limit *int, opStatus *uint32) ([]model.VideoInfoTab, error)
	UpdateVideoInfo(ctx context.Context, row model.VideoInfoTab) (err error)
	CountVideoInfo(ctx context.Context, mid int64, opStatus *uint32) (res int64, err error)
}
