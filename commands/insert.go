package commands

import (
	"regexp"
)

func Insert(source string) {
	gitURI := regexp.MustCompile("\\A(?:git|file|https?):\\/\\/|git@|.git\\z")
	githubShortcut := regexp.MustCompile("\\A[[:alnum:]-]+\\/[[:alnum:]-]+\\z")

	if githubShortcut.MatchString(source) {
		source = "https://github.com/" + source
	}

	if gitURI.MatchString(source) {
		runCommand("git clone", source)
	} else {
		exitWithMessage(1, source, "unknown")
	}
}
