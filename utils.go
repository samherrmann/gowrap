package main

import (
	"os"
	"path/filepath"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// currentFolderName returns the folder name
// of the current working directory
func currentFolderName() string {
	dir, err := os.Getwd()
	panicIf(err)
	return filepath.Base(dir)
}
