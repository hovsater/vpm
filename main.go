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
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "List installed plugins",
			Action: func(c *cli.Context) {
				commands.List()
			},
		},
		{
			Name:  "insert",
			Usage: "Insert plugin",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					commands.Insert(c.Args().First())
				}
			},
		},
		{
			Name:  "update",
			Usage: "Update installed plugins",
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					commands.UpdateAll()
				} else {
					commands.UpdateSingle(c.Args().First())
				}
			},
			BashComplete: pluginComplete,
		},
		{
			Name:  "remove",
			Usage: "Remove installed plugin",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					commands.Remove(c.Args().First())
				}
			},
			BashComplete: pluginComplete,
		},
	}

	app.Run(os.Args)
}

func pluginComplete(c *cli.Context) {
	if len(c.Args()) < 1 {
		commands.List()
	}

}
