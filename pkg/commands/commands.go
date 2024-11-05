package commands

import (
	"github.com/sqljames/factorio-mod-downloader/pkg/commands/download"
	"github.com/sqljames/factorio-mod-downloader/pkg/info"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func NewApp(logger *zap.Logger) *cli.App {

	app := &cli.App{
		Name:      info.GetApplicationName(),
		Usage:     info.Description,
		Authors:   info.Authors,
		Copyright: info.Copyright,
		Suggest:   true,
		Commands:  []*cli.Command{
			download.New(),
		},
	}
	if app.Metadata == nil {
		app.Metadata = make(map[string]interface{})
	}
	app.Metadata["logger"] = logger

	return app
}
