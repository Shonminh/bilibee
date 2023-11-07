package main

import (
	"github.com/Shonminh/bilibee/pkg/logger"
)

func main() {
	app, err := InitVideoTaskApp()
	if err != nil {
		logger.LogPanicf("InitVideoTaskApp error: %v", err)
	}
	logger.LogInfof("VideoTaskApp run...")
	app.Run()
}
