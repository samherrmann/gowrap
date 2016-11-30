package main

import "fmt"

var (
	// appName is the name of the
	// application to be built.
	appName = currentFolderName()

	// appVersion is the version of
	// the application to be built.
	appVersion = gitVersion()
)

func main() {
	c := defaultConfig()
	readOrSaveConfig(c)

	for _, target := range *c.Targets {
		goos, goarch := target.Parse()

		fmt.Println("Building " + buildName(appName, appVersion, goos, goarch) + "...")
		goGenerate()
		goBuild(appName, appVersion, goos, goarch)
	}
}
