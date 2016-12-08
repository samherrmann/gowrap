package main

import (
	"log"
	"path/filepath"

	"github.com/samherrmann/gowrap/gotools"
)

// runGoBuildChain executes the Go build tool-chain per configuration
func runGoBuildChain(t *[]gotools.Target) error {
	for _, target := range *t {
		goos, goarch := target.Parse()

		log.Println("Building " + buildName(appName, appVersion, goos, goarch) + "...")
		err := goGenerate()
		if err != nil {
			return err
		}

		err = goBuild(appName, appVersion, goos, goarch)
		if err != nil {
			return err
		}
	}
	return nil
}

// goGenerate executes the command "go generate"
func goGenerate() error {
	return gotools.Generate()
}

// goBuild executes the command "go build" for the desired
// target OS and architecture, and writes the generated
// executable to the 'outDir' directory.
func goBuild(name string, version string, goos string, goarch string) error {
	err := gotools.SetGoOS(goos)
	if err != nil {
		return err
	}

	err = gotools.SetGoArch(goarch)
	if err != nil {
		return err
	}

	path, err := buildPath(name, version, goos, goarch)
	if err != nil {
		return err
	}

	return gotools.Build("-o", path, "-ldflags", "-X main.version="+version)
}

// buildPath constructs a file path for a given target
func buildPath(name string, version string, os string, arch string) (string, error) {
	ext, err := gotools.ExeSuffix()
	if err != nil {
		return "", err
	}
	return filepath.Join(outputRoot, buildName(name, version, os, arch), name+ext), err
}

// buildName returns a build-name in the form of appname-version-os-arch
// ex: myapp-v1.0.0-linux-amd64
func buildName(name string, version string, os string, arch string) string {
	return name + "-" + version + "-" + os + "-" + arch
}
