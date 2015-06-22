package commands

import (
	"github.com/codegangsta/cli"
	"regexp"
)

func Insert(c *cli.Context) {
	plugin := c.Args().First()

	if plugin == "" {
		exitWithMessage(1, "no plugin provided")
	}

	if exists(plugin) {
		exitWithMessage(1, "already inserted")
	}

	gitURI := regexp.MustCompile("\\A(?:git|file|https?):\\/\\/|git@|.git\\z")
	githubShortcut := regexp.MustCompile("\\A[[:alnum:]-]+\\/[[:alnum:]-]+\\z")

	if githubShortcut.MatchString(plugin) {
		plugin = "https://github.com/" + plugin
	}

	if gitURI.MatchString(plugin) {
		runCommand("git clone", plugin)
	} else {
		exitWithMessage(1, plugin, "unknown")
	}
}
