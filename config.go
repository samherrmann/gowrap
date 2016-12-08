package main

import (
	"errors"

	"github.com/samherrmann/gowrap/gotools"
	"github.com/samherrmann/gowrap/jsonfile"
)

// NewConfig returns a Config struct.
func NewConfig() (*Config, error) {
	t := gotools.NewPlatform()

	c := &Config{
		Targets: &Targets{t.String()},
	}
	return c, nil
}

// Config represents the application configuration.
type Config struct {
	Targets *Targets `json:"targets"`
}

// Targets is a slice of strings representing the
// targeted platforms.
type Targets []string

// ToPlatforms convers the Targets slice to a gotools/Platform slice.
func (t *Targets) ToPlatforms() (*[]gotools.Platform, error) {
	ps := &[]gotools.Platform{}
	for _, target := range *t {
		p := gotools.NewPlatform()
		if err := p.Unmarshal(target); err != nil {
			return nil, err
		}
		*ps = append(*ps, *p)
	}
	return ps, nil
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
