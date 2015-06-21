package commands

import (
	"fmt"
	"path/filepath"
)

func List() {
	plugins, _ := filepath.Glob("*")
	for _, plugin := range plugins {
		fmt.Println(plugin)
	}
}
