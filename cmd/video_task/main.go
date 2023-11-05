package main

import (
	"github.com/Shonminh/bilibee/pkg/logger"
)

func main() {
	app, err := InitVideoTaskApp()
	if err != nil {
		logger.LogPanic(err)
	}
	logger.LogInfof("VideoTaskApp run...")
	app.Run()
}
