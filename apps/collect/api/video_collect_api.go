package api

import "context"

type VideoCollectService interface {
	CreateCronTask(ctx context.Context, mid int64) (err error)
	CollectVideoInfo(ctx context.Context) (err error)
}
