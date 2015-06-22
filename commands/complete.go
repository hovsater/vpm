package commands

import (
	"github.com/codegangsta/cli"
)

func PluginComplete(c *cli.Context) {
	if len(c.Args()) < 1 {
		List(c)
	}
}
