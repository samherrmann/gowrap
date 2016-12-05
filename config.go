package main

import (
	"errors"
	"strings"

	"github.com/samherrmann/gowrap/gotools"
	"github.com/samherrmann/gowrap/jsonfile"
)

func NewConfig() *Config {
	goos, err := gotools.GoOS()
	panicIf(err)

	goarch, err := gotools.GoArch()
	panicIf(err)

	t := &[]Target{
		Target(goos + "-" + goarch),
	}
	return &Config{
		Targets: t,
	}
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
