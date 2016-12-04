package git

import (
	osExec "os/exec"

	"github.com/samherrmann/gowrap/exec"
)

const (
	cmdName = "git"
)

// TagOfHEAD returns the tag of HEAD if one exists. If no tag exists, and empty
// string is returned.
func TagOfHEAD() (string, error) {
	if err := gitPath(); err != nil {
		return "", err
	}
	return exec.Command(cmdName, "tag", "--contains", "HEAD").OutputLine()
}

// HashOfHEAD returns the commit hash of HEAD.
func HashOfHEAD() (string, error) {
	if err := gitPath(); err != nil {
		return "", err
	}
	return exec.Command(cmdName, "rev-parse", "--short", "HEAD").OutputLine()
}

func gitPath() error {
	_, err := osExec.LookPath("git")
	return err
}
