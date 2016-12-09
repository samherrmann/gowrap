package main

import (
	"log"
	"path/filepath"

	"github.com/samherrmann/gowrap/gotools"
)

// runGoBuildChain executes the Go build tool-chain per configuration.
// If no errors are encountered, it returns the paths to the resulting
// executables.
func runGoBuildChain(platforms *[]gotools.Platform) (*[]string, error) {
	paths := &[]string{}

	for _, p := range *platforms {
		log.Println("Building " + buildName(appName, appVersion, &p) + "...")
		err := goGenerate()
		if err != nil {
			return nil, err
		}

		path, err := goBuild(appName, appVersion, &p)
		if err != nil {
			return nil, err
		}
		*paths = append(*paths, path)
	}
	return paths, nil
}

// goGenerate executes the command "go generate"
func goGenerate() error {
	return gotools.Generate()
}

// goBuild executes the command "go build" for the desired target OS and architecture.
// If no errors are encountered, it returns the path to the resulting executable.
func goBuild(name string, version string, p *gotools.Platform) (string, error) {
	err := gotools.SetGoOS(p.GOOS)
	if err != nil {
		return "", err
	}

	err = gotools.SetGoArch(p.GOARCH)
	if err != nil {
		return "", err
	}

	path, err := buildPath(name, version, p)
	if err != nil {
		return "", err
	}

	return path, gotools.Build("-o", path, "-ldflags", "-X main.version="+version)
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
