package main

import (
	"fmt"
	"os"
	"path/filepath"

	"strings"

	"github.com/samherrmann/gowrap/archive/gziptar"
	"github.com/samherrmann/gowrap/archive/zip"
)

// archiveBuilds copies the buid-artifacts into
// archive files for every distribution target
func archiveBuilds() {
	ap := assetPaths()

	fmt.Println("Archiving builds...")
	for _, bp := range buildPaths() {
		if isWindows(bp) {
			panicIf(zip.Make(filepath.Dir(bp), append(ap, bp)))

		} else {
			panicIf(gziptar.Make(filepath.Dir(bp), append(ap, bp)))
		}
	}
}

// buildPaths returns the paths of the built executables.
func buildPaths() []string {
	paths := []string{}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return err
	}

	panicIf(filepath.Walk(outputRoot, walkFunc))
	return paths
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
