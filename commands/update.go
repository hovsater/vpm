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
		updatePlugins([]string{plugin})
	}

}

func updatePlugins(plugins []string) {
	for _, plugin := range plugins {
		os.Chdir(plugin)
		runCommand("git pull origin master")
	}
}
