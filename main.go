package main

import "fmt"

var (
	// name of the application
	name = currentFolderName()

	// version of the application
	version = gitVersion()
)

func main() {
	c := defaultConfig()
	readOrSaveConfig(c)

	for _, target := range *c.Targets {
		goos, goarch := target.Parse()

		fmt.Println("Building " + buildName(name, version, goos, goarch) + "...")
		goGenerate()
		goBuild(name, version, goos, goarch)
	}
}
