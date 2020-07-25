package commands

import (
	"fmt"

	"github.com/tkw1536/ggman/constants"
	"github.com/tkw1536/ggman/program"
	"github.com/tkw1536/ggman/repos"
)

// CanonCommand is the entry point for the canon command
func CanonCommand(runtime *program.SubRuntime) (retval int, err string) {
	argc := runtime.Argc
	argv := runtime.Argv

	var lines []repos.CanLine
	var e error

	if argc == 2 {
		// if we have two argument, use the specific specification given
		lines = append(lines, repos.CanLine{Pattern: "", Canonical: argv[1]})
	} else {
		// else read the canon file
		lines, e = program.GetCanonOrPanic()
		if e != nil {
			err = constants.StringInvalidCanfile
			retval = constants.ErrorMissingConfig
			return
		}

	}

	// print the canonical url or error
	return printCanonOrError(lines, argv[0])
}

func printCanonOrError(lines []repos.CanLine, repo string) (retval int, err string) {
	// parse the repo uri
	uri, e := repos.NewRepoURI(repo)
	if e != nil {
		err = constants.StringUnparsedRepoName
		retval = constants.ErrorInvalidRepo
		return
	}

	// get the canonical one based on the canfile
	canonical := uri.CanonicalWith(lines)
	fmt.Println(canonical)

	return
}
