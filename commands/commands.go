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
	user, _ := user.Current()
	os.Chdir(filepath.Join(user.HomeDir, ".vim", "bundle"))
}

func exists(name string) {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		exitWithMessage(1, name, "not found")
	}
}

func plugins() (plugins []string) {
	plugins, _ = filepath.Glob("*")
	return
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
