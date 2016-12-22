package main

import (
	"log"

	"github.com/samherrmann/gowrap/app/config"
	"github.com/samherrmann/gowrap/app/gotools"
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
		log.Printf("Error getting folder name of current working directory: %v", err)
		return
	}

	appVersion, err = gitVersion()
	if err != nil {
		log.Printf("Error getting version from Git: %v", err)
		return
	}

	c := config.New()
	err = c.Load()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return
	}

	gt, err := gotools.New(appName, appVersion, *c.Targets)
	if err != nil {
		log.Printf("Error setting up build environment: %v", err)
		return
	}

	buildPaths, err := gt.RunBuildChain()
	if err != nil {
		log.Printf("Error running Go build chain: %v", err)
		return
	}

	err = archiveBuilds(buildPaths)
	if err != nil {
		log.Printf("Error archiving builds: %v", err)
		return
	}
}
