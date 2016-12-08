package gotools

import (
	"errors"
	"strings"
)

const (
	// platformStrNotationSep is the character that separates
	// the GOOS and GOARCH value in the platform string notation.
	platformStrNotationSep = "-"
)

// Platform represents a target operating system and compilation
// architecture. See the valid combinations of GOOS and GOARCH
// here: https://golang.org/doc/install/source#environment
type Platform struct {
	GOOS   string
	GOARCH string
}

// NewPlatform returns a new Platform type,
// initialized with the local system's
// operating system and architecture.
func NewPlatform() *Platform {
	return &Platform{
		GOOS:   GOHOSTOS,
		GOARCH: GOHOSTARCH,
	}
}

// Unmarshal parses the provided GOOS-GOARCH string and
// sets the values to the GOOS and GOARCH fields of the
// Platform struct.
func (p *Platform) Unmarshal(str string) error {
	sep := platformStrNotationSep

	if !strings.Contains(str, sep) {
		return errors.New("String " + str + " is not in the format of GOOS-GOARCH")
	}

	s := strings.Split(str, sep)
	s0 := s[0]
	s1 := s[1]

	for _, sp := range *SupportedPlatforms {
		if s0 == sp.GOOS && s1 == sp.GOARCH {
			p.GOOS = s0
			p.GOARCH = s1
			return nil
		}
	}
	return errors.New("The string does not represent a supported platform")
}

// String returns a Platform struct in string notation.
// i.e. GOOS-GOARCH
func (p *Platform) String() string {
	return p.GOOS + platformStrNotationSep + p.GOARCH
}
