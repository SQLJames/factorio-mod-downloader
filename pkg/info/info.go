package info

import (
	"github.com/urfave/cli/v2"
)

const (
	Description string = "A Cli tool to download factorio mods"
	Copyright   string = "Rhoat, LLC, 2024"
)

var (
	applicationName = "factorio-mod-downloader"
	Authors         = []*cli.Author{
		{
			Name:  "James Rhoat",
			Email: "James@Rhoat.com",
		},
	}
)

func GetApplicationName() string {
	return applicationName
}
