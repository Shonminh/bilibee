//go:build wireinject

package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
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

	wg := &sync.WaitGroup{}
	go app.runCollectVideoTask(ctx, wg)
	go app.runSyncVideoInfoToEsTask(ctx, wg)
	go app.runResetTaskStatusTask(ctx, wg)
	wg.Wait()
}

func (app *VideoTaskApp) runCollectVideoTask(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer func() {
		wg.Done()
		if err := recover(); err != nil {
			logger.LogErrorf("runCollectVideoTask panic, err=%s", err)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			logger.LogInfo("VideoTaskApp exit...")
			return
		default:
			if err := app.schema.CollectVideo(ctx); err != nil {
				logger.LogErrorf("CollectVideo failed, err=%s", err.Error())
			}
			logger.LogInfof("CollectVideo done...")
			time.Sleep(time.Second * 5)
			logger.LogInfo("CollectVideo sleep 5s...")
		}
	}
}

func (app *VideoTaskApp) runSyncVideoInfoToEsTask(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer func() {
		wg.Done()
		if err := recover(); err != nil {
			logger.LogErrorf("runSyncVideoInfoToEsTask panic, err=%s", err)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			logger.LogInfo("VideoTaskApp exit...")
			return
		default:
			if err := app.schema.SyncVideoInfoToEs(ctx); err != nil {
				logger.LogErrorf("SyncVideoInfoToEs failed, err=%s", err.Error())
			}
			logger.LogInfof("SyncVideoInfoToEs done...")
			time.Sleep(time.Second * 30)
			logger.LogInfo("SyncVideoInfoToEs sleep 30s...")
		}
	}
}

func (app *VideoTaskApp) runResetTaskStatusTask(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer func() {
		wg.Done()
		if err := recover(); err != nil {
			logger.LogErrorf("runResetTaskStatusTask panic, err=%s", err)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			logger.LogInfo("VideoTaskApp exit...")
			return
		default:
			if err := app.schema.ResetTaskUndoStatus(ctx); err != nil {
				logger.LogErrorf("ResetTaskStatus failed, err=%s", err.Error())
			}
			logger.LogInfof("ResetTaskStatus done...")
			time.Sleep(time.Hour)
			logger.LogInfo("ResetTaskStatus sleep 1 hour...")
		}
	}
}
