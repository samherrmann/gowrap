package main

import (
	"errors"
	"strings"

	"github.com/samherrmann/gowrap/gotools"
	"github.com/samherrmann/gowrap/jsonfile"
)

func NewConfig() (*Config, error) {
	goos, err := gotools.GoOS()
	if err != nil {
		return nil, err
	}

	goarch, err := gotools.GoArch()
	if err != nil {
		return nil, err
	}

	t := &[]Target{
		Target(goos + "-" + goarch),
	}
	c := &Config{
		Targets: t,
	}
	return c, nil
}

type Config struct {
	Targets *[]Target `json:"targets"`
}

type Target string

func (t *Target) Parse() (goos string, goarch string) {
	s := strings.Split(string(*t), "-")
	return s[0], s[1]
}

// readOrSave attempts to read the configuration file.
// If it cannot find the file, it will save a sample
// configuration file instead.
func readOrSaveConfig() (*Config, error) {
	filePath := "gowrap.json"

	c, err := NewConfig()
	if err != nil {
		return nil, err
	}

	err = jsonfile.Read(filePath, c)
	if err == nil {
		return c, nil
	}

	// re-initialize c since the read operation
	// above may have corrupted c.
	c, err = NewConfig()
	if err != nil {
		return nil, err
	}

	err = jsonfile.Write(filePath, c)
	if err != nil {
		return nil, err
	}

	return nil, errors.New("No 'gowrap.json' file found. " +
		"A sample file was created in the current directory. " +
		"Edit the file as required and re-run gowrap.")
}
