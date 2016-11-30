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
func goBuild(goos string, goarch string, name string) {
	os.Setenv("goos", goos)
	os.Setenv("goarch", goarch)

	out := filepath.Join(distRoot, name+exeSuffix())
	cmd("go", "build", "-o", out, "-ldflags", "-X main.version="+version).Run()
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
