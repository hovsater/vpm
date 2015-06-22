package main

import (
	"github.com/KevinSjoberg/vpm/commands"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "vpm"
	app.Usage = "Vim Plugin Manager"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "list",
			Usage:  "List plugins",
			Action: commands.List,
		},
		{
			Name:   "insert",
			Usage:  "Insert plugins",
			Action: commands.Insert,
		},
		{
			Name:         "update",
			Usage:        "Update installed plugins",
			Action:       commands.Update,
			BashComplete: commands.PluginComplete,
		},
		{
			Name:         "remove",
			Usage:        "Remove installed plugins",
			Action:       commands.Remove,
			BashComplete: commands.PluginComplete,
		},
	}

	app.Run(os.Args)
}
