package main

import (
	"fmt"
	"path/filepath"

	"github.com/samherrmann/gowrap/gotools"
)

// runGoBuildChain executes the Go build tool-chain per configuration
func runGoBuildChain(c *Config) {
	for _, target := range *c.Targets {
		goos, goarch := target.Parse()

		fmt.Println("Building " + buildName(appName, appVersion, goos, goarch) + "...")
		goGenerate()
		goBuild(appName, appVersion, goos, goarch)
	}
}

// goGenerate executes the command "go generate"
func goGenerate() {
	panicIf(gotools.Generate())
}

// goBuild executes the command "go build" for the desired
// target OS and architecture, and writes the generated
// executable to the 'outDir' directory.
func goBuild(name string, version string, goos string, goarch string) {
	panicIf(gotools.SetGoOS(goos))
	panicIf(gotools.SetGoArch(goarch))

	path := buildPath(name, version, goos, goarch)
	panicIf(gotools.Build("-o", path, "-ldflags", "-X main.version="+version))
}

// buildPath constructs a file path for a given target
func buildPath(name string, version string, os string, arch string) string {
	ext, err := gotools.ExeSuffix()
	panicIf(err)
	return filepath.Join(outputRoot, buildName(name, version, os, arch), name+ext)
}

// buildName returns a build-name in the form of appname-version-os-arch
// ex: myapp-v1.0.0-linux-amd64
func buildName(name string, version string, os string, arch string) string {
	return name + "-" + version + "-" + os + "-" + arch
}
