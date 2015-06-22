package commands

import (
	"github.com/codegangsta/cli"
	"os"
)

func Update(c *cli.Context) {
	plugin := c.Args().First()

	if plugin == "" {
		updatePlugins(plugins())
	} else {
		if !exists(plugin) {
			exitWithMessage(1, "not found")
		}

		updatePlugins([]string{plugin})
	}

}

func updatePlugins(plugins []string) {
	for _, plugin := range plugins {
		chdir()
		os.Chdir(plugin)
		runCommand("git pull origin master")
	}
}
