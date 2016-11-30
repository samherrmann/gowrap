package main

import (
	"os"
	"path/filepath"
)

func panicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// currentFolderName returns the folder name
// of the current working directory
func currentFolderName() string {
	dir, err := os.Getwd()
	panicIf(err)
	return filepath.Base(dir)
}

// buildName returns a build-name in the form of appname-version-os-arch
// ex: myapp-v1.0.0-linux-amd64
func buildName(name string, version string, os string, arch string) string {
	return name + "-" + version + "-" + os + "-" + arch
}
