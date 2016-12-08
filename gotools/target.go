package gotools

import "strings"

// Target represents a target operating system and compilation
// architecture. It is expected to be of the form GOOS-GOARCH.
// See the valid combinations of GOOS and GOARCH here:
// https://golang.org/doc/install/source#environment
type Target string

// NewTarget returns a new Target type.
func NewTarget(goos string, goarch string) Target {
	return Target(goos + "-" + goarch)
}

// Parse returns the GOOS and GOARCH value of the target.
func (t *Target) Parse() (goos string, goarch string) {
	s := strings.Split(string(*t), "-")
	return s[0], s[1]
}
