package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func List(c *cli.Context) {
	for _, plugin := range plugins() {
		fmt.Println(plugin)
	}
}
