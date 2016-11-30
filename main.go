package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	distRoot = "dist"
)

var (
	gooses   = &[]string{"windows", "linux"}
	goarches = &[]string{"amd64", "386"}
	name     = currentFolderName()
	version  = gitVersion()
)

func main() {
	goGenerate()

	for _, goos := range *gooses {
		for _, goarch := range *goarches {
			buildName := name + "-" + version + "-" + goos + "-" + goarch

			fmt.Println("Building " + buildName + "...")
			goBuild(goos, goarch, buildName)
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
