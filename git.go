package main

import "strings"

// gitVersion returns the tag of the HEAD, if one exists,
// or else the commit hash.
func gitVersion() string {
	out := cmd("git", "tag", "--contains", "HEAD").Output()
	tag := strings.Split(string(out), "\n")[0]

	if tag != "" {
		return tag
	}

	out = cmd("git", "rev-parse", "--short", "HEAD").Output()
	return strings.Split(string(out), "\n")[0]
}
