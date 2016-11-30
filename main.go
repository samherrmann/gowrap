package main

import (
	"fmt"
	"gowrap/config"
	"os"
)

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

func readOrSaveConfig(c *Config) {
	readErr, saveErr := config.ReadOrSave("gowrap.json", c)
	if readErr != nil && saveErr != nil {
		panicIf(saveErr)
	}
	if readErr != nil && saveErr == nil {
		fmt.Println("No 'gowrap.json' file found. " +
			"A sample file was created in the current directory. " +
			"Edit the file as required and re-run gowrap.")
		os.Exit(0)
	}
}
