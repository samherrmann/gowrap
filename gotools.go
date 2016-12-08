package main

import (
	"log"
	"path/filepath"

	"github.com/samherrmann/gowrap/gotools"
)

// runGoBuildChain executes the Go build tool-chain per configuration
func runGoBuildChain(platforms *[]gotools.Platform) error {
	for _, p := range *platforms {

		log.Println("Building " + buildName(appName, appVersion, &p) + "...")
		err := goGenerate()
		if err != nil {
			return err
		}

		err = goBuild(appName, appVersion, &p)
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
func goBuild(name string, version string, p *gotools.Platform) error {
	err := gotools.SetGoOS(p.GOOS)
	if err != nil {
		return err
	}

	err = gotools.SetGoArch(p.GOARCH)
	if err != nil {
		return err
	}

	path, err := buildPath(name, version, p)
	if err != nil {
		return err
	}

	return gotools.Build("-o", path, "-ldflags", "-X main.version="+version)
}

// buildPath constructs a file path for a given target
func buildPath(name string, version string, p *gotools.Platform) (string, error) {
	ext, err := gotools.ExeSuffix()
	if err != nil {
		return "", err
	}
	return filepath.Join(outputRoot, buildName(name, version, p), name+ext), err
}

// buildName returns a build-name in the form of appname-version-os-arch
// ex: myapp-v1.0.0-linux-amd64
func buildName(name string, version string, p *gotools.Platform) string {
	return name + "-" + version + "-" + p.GOOS + "-" + p.GOARCH
}
