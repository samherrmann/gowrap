package main

import "log"

const (
	// outputRoot is the output directory
	// for the build artifacts.
	outputRoot = "dist"
)

var (
	// appName is the name of the
	// application to be built.
	appName string

	// appVersion is the version of
	// the application to be built.
	appVersion = "latest"
)

func main() {
	var err error

	appName, err = currentFolderName()
	if err != nil {
		log.Println(err)
		return
	}

	appVersion, err = gitVersion()
	if err != nil {
		log.Println(err)
		return
	}

	config := NewConfig()
	err = config.Load()
	if err != nil {
		log.Println(err)
		return
	}

	targetPlatforms, err := config.Targets.ToPlatforms()
	if err != nil {
		log.Println(err)
		return
	}

	err = runGoBuildChain(targetPlatforms)
	if err != nil {
		log.Println(err)
		return
	}

	err = archiveBuilds()
	if err != nil {
		log.Println(err)
		return
	}
}
