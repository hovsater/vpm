package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"path/filepath"
)

func List(c *cli.Context) {
	plugins, _ := filepath.Glob("*")
	for _, plugin := range plugins {
		fmt.Println(plugin)
	}
}
