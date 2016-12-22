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
	gotools *GoTools

	goos   string
	goarch string
}

// Platform returns a new Platform type,
// initialized with the local system's
// operating system and architecture.
func (gt *GoTools) Platform() (*Platform, error) {
	goos, err := gt.GOHOSTOS()
	if err != nil {
		return nil, err
	}
	goarch, err := gt.GOHOSTARCH()
	if err != nil {
		return nil, err
	}

	p := &Platform{
		gotools: gt,
		goos:    goos,
		goarch:  goarch,
	}
	return p, nil
}

// GOOS returns the value of GOOS for this Platform struct.
func (p *Platform) GOOS() string {
	return p.goos
}

// SetGOOS sets the value for GOOS on this Platform struct.
func (p *Platform) SetGOOS(v string) {
	p.goos = v
}

// GOARCH returns the value of GOARCH for this Platform struct.
func (p *Platform) GOARCH() string {
	return p.goarch
}

// SetGOARCH sets the value for GOARCH on this Platform struct.
func (p *Platform) SetGOARCH(v string) {
	p.goarch = v
}

// Unmarshal parses the provided GOOS-GOARCH string and
// sets the values to the GOOS and GOARCH fields of the
// Platform struct.
func (p *Platform) Unmarshal(str string) error {
	sep := platformStrNotationSep

	if !strings.Contains(str, sep) {
		return errors.New("String \"" + str + "\" is not in the format of GOOS-GOARCH")
	}

	s := strings.Split(str, sep)
	s0 := s[0]
	s1 := s[1]

	sp, err := p.gotools.SupportedPlatforms()
	if err != nil {
		return err
	}

	for _, sp := range *sp {
		if s0 == sp.goos && s1 == sp.goarch {
			p.goos = s0
			p.goarch = s1
			return nil
		}
	}
	return errors.New("The string \"" + str + "\" does not represent a supported platform")
}

// String returns a Platform struct in string notation
// in the format of GOOS-GOARCH.
func (p *Platform) String() string {
	return p.goos + platformStrNotationSep + p.goarch
}
