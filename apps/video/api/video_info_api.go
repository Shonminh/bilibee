package api

import (
	"context"
	"github.com/Shonminh/bilibee/apps/video/internal/repository/model"
)

type VideoInfoService interface {
	CreateCronTask(ctx context.Context, mid int64, taskType model.TaskType) (err error)
	CollectVideoInfo(ctx context.Context) (err error)
	SyncVideoInfoToEs(ctx context.Context) (err error)
	ResetTaskUndoStatus(ctx context.Context) (err error)
}
