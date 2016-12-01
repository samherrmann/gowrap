package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/samherrmann/gowrap/jsonfile"
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

// readOrSave attempts to read the configuration file.
// If it cannot find the file, it will save a sample
// configuration file instead.
func readOrSaveConfig() (*Config, error) {
	c := NewConfig()
	filePath := "gowrap.json"

	if err := jsonfile.Read(filePath, c); err == nil {
		return c, nil
	}

	c = NewConfig()
	err := jsonfile.Write(filePath, c)
	panicIf(err)

	return nil, errors.New("No 'gowrap.json' file found. " +
		"A sample file was created in the current directory. " +
		"Edit the file as required and re-run gowrap.")
}

// runGoBuildChain executes the Go build tool-chain per configuration
func runGoBuildChain(c *Config) {
	for _, target := range *c.Targets {
		goos, goarch := target.Parse()

		fmt.Println("Building " + buildName(appName, appVersion, goos, goarch) + "...")
		goGenerate()
		goBuild(appName, appVersion, goos, goarch)
	}
}
