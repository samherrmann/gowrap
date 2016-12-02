package main

// gitVersion returns the tag of the HEAD, if one exists,
// or else the commit hash.
func gitVersion() string {
	tag := cmd("git", "tag", "--contains", "HEAD").OutputLine()

	if tag != "" {
		return tag
	}

	return cmd("git", "rev-parse", "--short", "HEAD").OutputLine()
}
