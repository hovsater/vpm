package commands

import "os"

func Remove(name string) {
	exists(name)
	os.RemoveAll(name)
	exitWithMessage(0, name, "removed")
}
