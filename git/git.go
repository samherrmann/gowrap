package git

import (
	"os/exec"
	"strings"
)

const (
	cmdName = "git"
)

// TagOfHEAD returns the tag of HEAD if one exists. If no tag exists, and empty
// string is returned.
func TagOfHEAD() (string, error) {
	return cmdOutput(cmdName, "tag", "--contains", "HEAD")
}

// HashOfHEAD returns the commit hash of HEAD.
func HashOfHEAD() (string, error) {
	return cmdOutput(cmdName, "rev-parse", "--short", "HEAD")
}

// cmdOutput executes the command specified by name and
// returns the first line of the standard output.
func cmdOutput(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).Output()
	return strings.TrimRight(string(out), "\n"), err
}
