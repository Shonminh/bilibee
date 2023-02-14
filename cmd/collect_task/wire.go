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

	"github.com/Shonminh/bilibee/apps/collect/access/crontask"
	collect "github.com/Shonminh/bilibee/apps/collect/wire"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/logger"
)

type CollectTaskApp struct {
	db     *gorm.DB
	schema *crontask.VideoCollectTaskSchema
}

func InitCollectTaskApp() (*CollectTaskApp, error) {
	wire.Build(
		db.NewDB,
		collect.CollectTaskSet,
		wire.Struct(new(CollectTaskApp), "*"),
	)

	return &CollectTaskApp{}, nil
}

func (app *CollectTaskApp) Run() {

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 绑定db 连接
	ctx := db.BindDbContext(context.Background(), app.db)

	for {
		select {
		case <-sigChan:
			logger.LogInfo("CollectTaskApp exit")
			return
		default:
			if err := app.schema.CollectVideo(ctx); err != nil {
				logger.LogErrorf("CollectVideo failed, err=%s", err.Error())
			}
			time.Sleep(time.Second * 5)
		}
	}
}
