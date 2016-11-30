package main

import (
	"fmt"
	"gowrap/config"
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

// readOrSave attempts to read the configuration file. If it
// cannot find the file, it will save a sample configuration
// file instead.
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
