package commands

import (
	"fmt"
	"os"
)

func Remove(name string) {
	exists(name)
	os.RemoveAll(name)
	fmt.Printf("vpm: %s removed\n", name)
}
