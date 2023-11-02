package main

import (
	"github.com/Shonminh/bilibee/pkg/logger"
)

func main() {
	app, err := InitCollectTaskApp()
	if err != nil {
		logger.LogPanic(err)
	}
	logger.LogInfof("CollectTaskApp run...")
	app.Run()
}
