package main

import (
	"os"
	"path/filepath"
)

// currentFolderName returns the folder name
// of the current working directory
func currentFolderName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Base(dir), err
}
