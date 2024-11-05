package main

import (
	"github.com/sqljames/factorio-mod-downloader/pkg/application"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	application.Run(logger)
}
