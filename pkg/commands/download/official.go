package download

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"time"

	"io"

	"github.com/sqljames/factorio-mod-downloader/pkg/commands/semver"
	"github.com/sqljames/factorio-mod-downloader/pkg/flags"
	"github.com/sqljames/factorio-mod-downloader/pkg/net"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func official() *cli.Command {
	return &cli.Command{
		Name:    "official",
		Usage:   "downloads an official mod from the mod portal",
		Action:  actionDownloadOfficial,
		Aliases: []string{"o"},
		Flags: []cli.Flag{
			flags.NameOfMod,
			flags.DownloadDestination,
			flags.ModVersion,
			flags.FactorioVersion,
			flags.Username,
			flags.APIToken,
		},
	}
}

const (
	baseurl = "https://mods.factorio.com"
)

var (
	modCall     = "%s/api/mods/%s?username=%s&token=%s"
	modDownload = "%s/%s?username=%s&token=%s"
)

func actionDownloadOfficial(cliContext *cli.Context) error {
	logger := cliContext.App.Metadata["logger"].(*zap.Logger)
	modName := cliContext.String(flags.NameFlagName)
	destination := cliContext.Path(flags.Destination)
	version := cliContext.String(flags.Version)
	factorioVersion, err := semver.GetMajorMinor(cliContext.String(flags.GameVersion))
	if err != nil {
		return err
	}
	user := cliContext.String(flags.User)
	token := cliContext.String(flags.Token)
	logger.Info("Download Official Params", zap.String("Name", modName),
		zap.String("Name", modName),
		zap.String("destination", destination),
		zap.String("version", version),
		zap.String("factorioVersion", factorioVersion),
		zap.String("user", user),
		zap.String("token", token))
	url := fmt.Sprintf(modCall, baseurl, modName, user, token)
	logger.Info("built url", zap.String("url", url))
	resp, err := net.Get(logger, url)
	if err != nil {
		logger.Error("failed to get mod url data", zap.String("URL", url), zap.Error(err))
		return fmt.Errorf("failed to get mod url data from %s: %v", url, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body", zap.Error(err))
		return fmt.Errorf("failed to read response body: %v", err)
	}
	var modPortalReleases modPortalResponse
	modPortalReleases.Unmarshal(data)

	logger.Debug("modPortalList", zap.Array("results", zapcore.ArrayMarshalerFunc(
		func(ae zapcore.ArrayEncoder) error {
			for _, v := range modPortalReleases.Releases {
				ae.AppendString(v.FileName)
				ae.AppendString(v.Version)
				ae.AppendString(v.InfoJSON.FactorioVersion)
			}
			return nil
		},
	)))
	filters := []FilterFunc{}

	if version != "" {
		logger.Info("Version supplied, attempting to pull that version.")
		filters = append(filters, FilterByModVersion(version))
	}

	filters = append(filters, FilterByFactoioVersion(factorioVersion))
	filteredRelease, err := FilterReleases(logger, modPortalReleases.Releases, filters...)
	if err != nil {
		logger.Error("Error filtering releases", zap.Error(err))
		return err
	}
	logger.Info("Viable Release", zap.String("file", filteredRelease.FileName))
	filePath := filepath.Join(destination, filteredRelease.FileName)
	file, err := os.Create(filePath)
	if err != nil {
		logger.Error("failed to create file", zap.String("filePath", filePath), zap.Error(err))
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer file.Close()
	err = net.Zip(logger, file, fmt.Sprintf(modDownload, baseurl, filteredRelease.DownloadURL, user, token))
	if err != nil {
		logger.Error("Failed to download file", zap.String("filePath", filePath), zap.Error(err))
		return fmt.Errorf("Failed to download file: %v", err)
	}
	return nil
}

type modPortalResponse struct {
	Category          string    `json:"category"`
	DownloadsCount    int       `json:"downloads_count"`
	LastHighlightedAt string    `json:"last_highlighted_at"`
	Name              string    `json:"name"`
	Owner             string    `json:"owner"`
	Releases          []Release `json:"releases"`
	Score             float64   `json:"score"`
	Summary           string    `json:"summary"`
	Thumbnail         string    `json:"thumbnail"`
	Title             string    `json:"title"`
}
type InfoJSON struct {
	FactorioVersion string `json:"factorio_version"`
}
type Release struct {
	DownloadURL string    `json:"download_url"`
	FileName    string    `json:"file_name"`
	InfoJSON    InfoJSON  `json:"info_json"`
	ReleasedAt  time.Time `json:"released_at"`
	Sha1        string    `json:"sha1"`
	Version     string    `json:"version"`
}

func (mpr *modPortalResponse) Unmarshal(data []byte) {
	json.Unmarshal(data, &mpr)
}

type FilterFunc func(Release) bool

func FilterReleases(logger *zap.Logger, releases []Release, filters ...FilterFunc) (Release, error) {
	var result []Release
	for _, release := range releases {
		match := true
		for i, filter := range filters {
			if !filter(release) {
				logger.Debug("Filter failed", zap.Int("filter_index", i), zap.String("mod", release.FileName))
				match = false
				break // Stop further checks for this release
			} else {
				logger.Debug("Filter passed", zap.Int("filter_index", i), zap.String("mod", release.FileName))
			}
		}
		if match {
			logger.Debug("Valid match", zap.String("mod", release.FileName))
			result = append(result, release)
		}
	}

	logger.Debug("Filtered results", zap.Array("results", zapcore.ArrayMarshalerFunc(
		func(ae zapcore.ArrayEncoder) error {
			for _, v := range result {
				ae.AppendObject(zapcore.ObjectMarshalerFunc(func(oe zapcore.ObjectEncoder) error {
					oe.AddString("FileName", v.FileName)
					oe.AddString("Version", v.Version)
					oe.AddString("FactorioVersion", v.InfoJSON.FactorioVersion)
					return nil
				}))
			}
			return nil
		},
	)))

	if len(result) == 0 {
		return Release{}, fmt.Errorf("no matching releases found")
	}

	if len(result) > 1 {
		sort.Slice(result, func(i, j int) bool {
			return result[i].ReleasedAt.After(result[j].ReleasedAt)
		})
	}

	return result[0], nil
}

func FilterByFactoioVersion(version string) FilterFunc {
	return func(r Release) bool {
		return fmt.Sprintf("v%s", r.InfoJSON.FactorioVersion) == version
	}
}

func FilterByModVersion(version string) FilterFunc {
	return func(r Release) bool {
		return r.Version == version
	}
}
