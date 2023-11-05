//go:build wireinject

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/Shonminh/bilibee/apps/video/access/crontask"
	collect "github.com/Shonminh/bilibee/apps/video/wire"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/logger"
)

type VideoTaskApp struct {
	db     *gorm.DB
	schema *crontask.VideoCollectTaskSchema
}

func InitVideoTaskApp() (*VideoTaskApp, error) {
	wire.Build(
		db.NewDB,
		collect.CollectTaskSet,
		wire.Struct(new(VideoTaskApp), "*"),
	)

	return &VideoTaskApp{}, nil
}

func (app *VideoTaskApp) Run() {

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 绑定db 连接
	ctx := db.BindDbContext(context.Background(), app.db)

	for {
		select {
		case <-sigChan:
			logger.LogInfo("VideoTaskApp exit")
			return
		default:
			if err := app.schema.CollectVideo(ctx); err != nil {
				logger.LogErrorf("CollectVideo failed, err=%s", err.Error())
			}
			time.Sleep(time.Second * 5)
			logger.LogInfo("CollectVideo sleep 5s...")
		}
	}
}
