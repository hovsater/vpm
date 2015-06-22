package commands

import (
	"github.com/codegangsta/cli"
	"os"
)

func Remove(c *cli.Context) {
	plugin := c.Args().First()

	if plugin == "" {
		exitWithMessage(1, "no plugin provided")
	}

	if !exists(plugin) {
		exitWithMessage(1, "not found")
	}

	os.RemoveAll(plugin)
	exitWithMessage(0, plugin, "removed")
}
