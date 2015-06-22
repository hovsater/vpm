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

	exists(plugin)
	os.RemoveAll(plugin)
	exitWithMessage(0, plugin, "removed")
}
