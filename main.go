package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	// targets for which to build
	// TODO: Get list from config file.
	gooses   = &[]string{"windows", "linux"}
	goarches = &[]string{"amd64", "386"}

	// name of the application
	name = currentFolderName()

	// version of the application
	version = gitVersion()
)

func main() {
	goGenerate()

	for _, goos := range *gooses {
		for _, goarch := range *goarches {
			fmt.Println("Building " + buildName(name, version, goos, goarch) + "...")
			goBuild(name, version, goos, goarch)
		}
	}
}

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
