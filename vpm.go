package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	user, _ := user.Current()
	os.Chdir(filepath.Join(user.HomeDir, ".vim", "bundle"))

	app := cli.NewApp()
	app.Name = "vpm"
	app.Usage = "Vim Plugin Manager"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "List installed plugins",
			Action: func(c *cli.Context) {
				List()
			},
		},
		{
			Name:  "insert",
			Usage: "Insert plugin",
			Action: func(c *cli.Context) {
				Insert(c.Args().First())
			},
		},
		{
			Name:  "update",
			Usage: "Update installed plugins",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					UpdateOne(c.Args().First())
				} else {
					UpdateAll()
				}
			},
			BashComplete: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					return
				}

				List()
			},
		},
		{
			Name:  "remove",
			Usage: "Remove installed plugin",
			Action: func(c *cli.Context) {
				Remove(c.Args().First())
			},
			BashComplete: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					return
				}

				List()
			},
		},
	}

	app.Run(os.Args)
}

func List() {
	plugins, _ := filepath.Glob("*")
	for _, plugin := range plugins {
		fmt.Println(plugin)
	}
}

func Insert(source string) {
	gitURI := regexp.MustCompile("\\A(?:git|file|https?):\\/\\/|git@|.git\\z")
	githubShortcut := regexp.MustCompile("\\A[[:alnum:]-]+\\/[[:alnum:]-]+\\z")

	if githubShortcut.MatchString(source) {
		source = "https://github.com/" + source
	}

	if gitURI.MatchString(source) {
		runCommand("git clone", source)
	} else {
		fmt.Printf("vpm: %s can not be inserted\n", source)
		os.Exit(1)
	}
}

func Update(plugins []string) {
	for _, plugin := range plugins {
		os.Chdir(plugin)
		runCommand("git pull origin master")
	}
}

func UpdateOne(name string) {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		fmt.Printf("vpm: %s not found\n", name)
		os.Exit(1)
	}

	Update([]string{name})
}

func UpdateAll() {
	plugins, _ := filepath.Glob("*")
	Update(plugins)
}

func Remove(name string) {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		fmt.Printf("vpm: %s not found\n", name)
		os.Exit(1)
	}

	os.RemoveAll(name)
	fmt.Printf("vpm: %s removed\n", name)
}

func runCommand(cmd ...string) {
	command := exec.Command("sh", "-c", strings.Join(cmd, " "))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
