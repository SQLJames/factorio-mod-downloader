package download

import (
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "download",
		Usage:   "Allows users to download mods",
		Aliases: []string{"d"},
		Subcommands: []*cli.Command{
			unofficial(),
			official(),
		},
	}
}
