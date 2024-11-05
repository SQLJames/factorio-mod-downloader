package flags

import "github.com/urfave/cli/v2"

const (
	NameFlagName string = "name"
	URL          string = "url"
	Destination  string = "destination"
	Version      string = "version"
	GameVersion  string = "factorioVersion"
	Token        string = "token"
	User         string = "user"
)

var (
	APIToken *cli.StringFlag = &cli.StringFlag{
		Name:     Token,
		Usage:    "Token for the factorio user.",
		Required: true,
	}
	Username *cli.StringFlag = &cli.StringFlag{
		Name:     User,
		Usage:    "Name of the user.",
		Required: true,
	}
	NameOfMod *cli.StringFlag = &cli.StringFlag{
		Name:     NameFlagName,
		Usage:    "Name Of the mod you are downloading. Suggested Name_version",
		Required: true,
	}
	ModVersion *cli.StringFlag = &cli.StringFlag{
		Name:     Version,
		Usage:    "Version of the mod you want to download",
		Required: false,
	}
	FactorioVersion *cli.StringFlag = &cli.StringFlag{
		Name:     GameVersion,
		Usage:    "Version of factorio you are running",
		Required: true,
	}
	DownloadURL *cli.StringFlag = &cli.StringFlag{
		Name:     URL,
		Usage:    "URL that you can download the mod zip package from. E.G. https://github.com/Suprcheese/Squeak-Through/releases/download/1.9.0/Squeak.Through_1.9.0.zip",
		Required: true,
	}
	DownloadDestination *cli.PathFlag = &cli.PathFlag{
		Name:     Destination,
		Usage:    "Destination where the mods will be downloaded.",
		Required: true,
	}
)
