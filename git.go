package main

import "github.com/samherrmann/gowrap/git"

// gitVersion returns one of the following:
// - Git tag of the HEAD if one exists, or
// - Commit hash of the HEAD, or
// - Empty string if Git is not in the PATH.
func gitVersion() (string, error) {
	tag, err := git.TagOfHEAD()
	if err != nil {
		return "", err
	}

	if tag != "" {
		return tag, nil
	}

	hash, err := git.HashOfHEAD()
	return hash, err
}
