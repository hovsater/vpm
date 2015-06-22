package commands

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func init() {
	chdir()
}

func chdir() {
	user, _ := user.Current()
	os.Chdir(filepath.Join(user.HomeDir, ".vim", "bundle"))
}

func exists(name string) bool {
	user, _ := user.Current()
	name = filepath.Join(user.HomeDir, ".vim", "bundle", filepath.Base(name))
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func plugins() []string {
	plugins, _ := filepath.Glob("*")
	return plugins
}

func runCommand(cmd ...string) {
	command := exec.Command("sh", "-c", strings.Join(cmd, " "))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}

func exitWithMessage(exitCode int, text ...string) {
	fmt.Printf("vpm: %s", text[0])
	for _, t := range text[1:] {
		fmt.Printf(" %s", t)
	}
	fmt.Printf("\n")
	os.Exit(exitCode)
}
