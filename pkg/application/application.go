package application

import (
	"os"

	"github.com/sqljames/factorio-mod-downloader/pkg/commands"
	"go.uber.org/zap"
)

func Run(logger *zap.Logger) {

	// Build command tree.
	cmd := commands.NewApp(logger)

	// Execute.
	if err := cmd.Run(os.Args); err != nil {
		logger.Fatal("Error running command", zap.Error(err))
	}
}
