package main

import "github.com/samherrmann/gowrap/git"

// gitVersion returns one of the following:
// - Git tag of the HEAD if one exists, or
// - Commit hash of the HEAD, or
// - Empty string if Git is not in the PATH.
func gitVersion() string {
	tag, err := git.TagOfHEAD()
	panicIf(err)
	if tag != "" {
		return tag
	}

	hash, err := git.HashOfHEAD()
	panicIf(err)
	return hash
}
