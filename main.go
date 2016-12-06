package main

import "fmt"

var (
	// version is the version of this application.
	version = "latest"

	// appName is the name of the
	// application to be built.
	appName = currentFolderName()

	// appVersion is the version of
	// the application to be built.
	appVersion = gitVersion()

	outputRoot = "dist"
)

func main() {
	fmt.Println("Version: " + version)

	config, err := readOrSaveConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	runGoBuildChain(config)
	archiveBuilds()
}
