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
		fmt.Printf("vpm: %s not found\n", name)
		os.Exit(1)
	}
}

func runCommand(cmd ...string) {
	command := exec.Command("sh", "-c", strings.Join(cmd, " "))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
