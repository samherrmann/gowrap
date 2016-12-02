package main

import (
	"os"
	"path/filepath"
)

// goGenerate executes the command "go generate"
func goGenerate() {
	cmd("go", "generate").Run()
}

// goBuild executes the command "go build" for the desired
// target OS and architecture, and writes the generated
// executable to the 'outDir' directory.
func goBuild(name string, version string, goos string, goarch string) {
	os.Setenv("goos", goos)
	os.Setenv("goarch", goarch)

	out := distPath(name, version, goos, goarch)
	cmd("go", "build", "-o", out, "-ldflags", "-X main.version="+version).Run()
}

// distPath constructs a file path for a given target
func distPath(name string, version string, os string, arch string) string {
	return filepath.Join("dist", buildName(name, version, os, arch), name+exeSuffix())
}

// exeSuffix returns ".exe" if the GOOS
// environment variable is set to
// "windows".
func exeSuffix() string {
	if os.Getenv("GOOS") == "windows" {
		return ".exe"
	}
	return ""
}

// goOS returns the value of GOOS
func goOS() string {
	return cmd("go", "env", "GOOS").OutputLine()
}

// goArch returns the value of GOARCH
func goArch() string {
	return cmd("go", "env", "GOARCH").OutputLine()
}
