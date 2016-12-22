// Package gotools is an application level wrapper package for the github.com/samherrmann/gowrap/gotools package.
package gotools

import (
	"log"
	"path/filepath"

	"github.com/samherrmann/gowrap/gotools"
)

// GoTools represents a Go build environment.
type GoTools struct {
	name      string
	version   string
	platforms *[]Platform
	gotools   *gotools.GoTools
}

// New returns a GoTools instance.
func New(name string, version string, platforms []string) (*GoTools, error) {
	gt := gotools.New()
	ps := &[]Platform{}

	for _, pStr := range platforms {
		p, err := gt.Platform()
		if err != nil {
			return nil, err
		}
		err = p.Unmarshal(pStr)
		if err != nil {
			return nil, err
		}
		*ps = append(*ps, p)
	}

	return &GoTools{
		name:      name,
		version:   version,
		platforms: ps,
		gotools:   gt,
	}, nil
}

// RunBuildChain executes the Go build tool-chain per configuration. If no errors are encountered,
// it returns the paths to the resulting executables.
func (gt *GoTools) RunBuildChain() (*[]string, error) {
	paths := &[]string{}

	// go generate...
	err := gt.generate()
	if err != nil {
		return nil, err
	}

	// go build...
	for _, p := range *gt.platforms {
		path, err := gt.build(p)
		if err != nil {
			return nil, err
		}
		*paths = append(*paths, path)
	}
	return paths, nil
}

// generate executes the command "go generate"
func (gt *GoTools) generate() error {
	log.Println("Running go generate...")

	return gt.gotools.Generate()
}

// build executes the command "go build" for the desired target OS and architecture.
// If no errors are encountered, it returns the path to the resulting executable.
func (gt *GoTools) build(p Platform) (string, error) {
	log.Println("Building " + gt.buildName(p) + "...")

	err := gt.gotools.SetGOOS(p.GOOS())
	if err != nil {
		return "", err
	}

	err = gt.gotools.SetGOARCH(p.GOARCH())
	if err != nil {
		return "", err
	}

	path, err := gt.buildPath(p)
	if err != nil {
		return "", err
	}
	return path, gt.gotools.Build("-o", path, "-ldflags", "-X main.version="+gt.version)
}

// buildPath constructs a file path for a given target
func (gt *GoTools) buildPath(p Platform) (string, error) {
	ext, err := gt.gotools.ExeSuffix()
	if err != nil {
		return "", err
	}
	return filepath.Join("dist", gt.buildName(p), gt.name+ext), err
}

// buildName returns a build-name in the form of appname-version-os-arch
// ex: myapp-v1.0.0-linux-amd64
func (gt *GoTools) buildName(p Platform) string {
	return gt.name + "-" + gt.version + "-" + p.GOOS() + "-" + p.GOARCH()
}

// Platform represents a target operating system.
type Platform interface {
	GOOS() string
	GOARCH() string
}
