package commands

import (
	"os"
	"path/filepath"
)

func Update(plugins []string) {
	for _, plugin := range plugins {
		os.Chdir(plugin)
		runCommand("git pull origin master")
	}
}

func UpdateAll() {
	plugins, _ := filepath.Glob("*")
	Update(plugins)
}

func UpdateSingle(name string) {
	exists(name)
	Update([]string{name})
}
