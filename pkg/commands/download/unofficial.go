package download

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sqljames/factorio-mod-downloader/pkg/flags"
	"github.com/sqljames/factorio-mod-downloader/pkg/net"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func unofficial() *cli.Command {
	return &cli.Command{
		Name:    "unofficial",
		Usage:   "downloads an unofficial mod from alternate sources uring a web request",
		Action:  actionDownloadUnofficial,
		Aliases: []string{"u"},
		Flags: []cli.Flag{
			flags.DownloadURL,
			flags.NameOfMod,
			flags.DownloadDestination,
		},
	}
}

func actionDownloadUnofficial(cliContext *cli.Context) error {
	logger := cliContext.App.Metadata["logger"].(*zap.Logger)
	modName := cliContext.String(flags.NameFlagName)
	url := cliContext.String(flags.URL)
	destination := cliContext.Path(flags.Destination)
	filePath := filepath.Join(filepath.Clean(destination), filepath.Clean(modName))

	file, err := os.Create(filePath)
	if err != nil {
		logger.Error("failed to create file", zap.String("filePath", filePath), zap.Error(err))
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer file.Close()
	err = net.Zip(logger, file, url)
	if err != nil {
		logger.Error("Failed to download file", zap.String("filePath", filePath), zap.Error(err))
		return fmt.Errorf("Failed to download file: %v", err)
	}

	logger.Info("Mod Downloaded", zap.String("name", modName), zap.String("path", filePath))
	return nil
}
