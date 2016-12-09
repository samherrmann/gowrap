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
		log.Printf("Error while getting folder name of current working directory: %v", err)
		return
	}

	appVersion, err = gitVersion()
	if err != nil {
		log.Printf("Error while getting version from Git: %v", err)
		return
	}

	config := NewConfig()
	err = config.Load()
	if err != nil {
		log.Printf("Error while loading config: %v", err)
		return
	}

	buildPaths, err := runGoBuildChain(config.Platforms())
	if err != nil {
		log.Printf("Error while running Go build chain: %v", err)
		return
	}

	err = archiveBuilds(buildPaths)
	if err != nil {
		log.Printf("Error while archiving builds: %v", err)
		return
	}
}
