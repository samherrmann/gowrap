package main

import "os/exec"

// gitVersion returns one of the following:
// - Git tag of the HEAD if one exists, or
// - Commit hash of the HEAD, or
// - Empty string if Git is not in the PATH.
func gitVersion() string {
	cmdName := "git"

	if _, err := exec.LookPath(cmdName); err != nil {
		return ""
	}

	tag := cmd(cmdName, "tag", "--contains", "HEAD").OutputLine()
	if tag != "" {
		return tag
	}

	return cmd(cmdName, "rev-parse", "--short", "HEAD").OutputLine()
}
