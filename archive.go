package main

import (
	"log"
	"os"
	"path/filepath"

	"strings"

	"github.com/samherrmann/gowrap/archive/gziptar"
	"github.com/samherrmann/gowrap/archive/zip"
)

// archiveBuilds copies the buid-artifacts into
// archive files for every distribution target
func archiveBuilds(buildPaths *[]string) error {
	log.Println("Archiving builds...")

	aps := assetPaths()

	for _, bp := range *buildPaths {
		if isWindows(bp) {
			err := zip.Make(filepath.Dir(bp), append(aps, bp))
			if err != nil {
				return err
			}

		} else {
			err := gziptar.Make(filepath.Dir(bp), append(aps, bp))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// assetPaths returns the paths of asset files,
// i.e. README.md and LICENSE, if they exist.
func assetPaths() []string {
	files := []string{
		"README.md",
		"LICENSE",
	}

	for _, f := range files {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			f = files[len(files)-1]
			files = files[:len(files)-1]
		}
	}
	return files
}

// isWindows returns true if the provided 'buildName'
// contains the substring 'windows'
func isWindows(buildName string) bool {
	return strings.Contains(buildName, "windows")
}
