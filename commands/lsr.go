package commands

import (
	"fmt"

	"github.com/tkw1536/ggman/constants"
	"github.com/tkw1536/ggman/git"
	"github.com/tkw1536/ggman/program"
	"github.com/tkw1536/ggman/repos"
)

// LSRCommand is the entry point for the lsr command
func LSRCommand(runtime *program.SubRuntime) (retval int, err string) {
	shouldCanon := runtime.Flag

	// conditionally read the canon lines
	var lines []repos.CanLine
	var e error
	if shouldCanon {
		lines, e = program.GetCanonOrPanic()
		if e != nil {
			err = constants.StringInvalidCanfile
			retval = constants.ErrorMissingConfig
			return
		}
	}

	// get the root
	root := runtime.Root

	// find all the repos
	rs := repos.Repos(root, runtime.For)

	// and print them
	for _, repo := range rs {
		remote, err := git.Default.GetRemote(repo)
		if err == nil {
			if shouldCanon {
				printCanonOrError(lines, remote)
			} else {
				fmt.Println(remote)

			}
		}
	}

	return
}
