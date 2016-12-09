package main

import (
	"errors"

	"os"

	"github.com/samherrmann/gowrap/gotools"
	"github.com/samherrmann/gowrap/jsonfile"
)

const (
	configFilePath = "gowrap.json"
)

// NewConfig returns a Config struct.
func NewConfig() *Config {
	t := gotools.NewPlatform()

	c := &Config{
		Targets: &Targets{t.String()},
	}
	return c
}

// Config represents the application configuration.
type Config struct {
	Targets *Targets `json:"targets"`
}

// Save writes the Config struct to the
// JSON configuration file.
func (c *Config) Save() error {
	return jsonfile.Write(configFilePath, c)
}

// Load decodes the JSON configuration file into
// the Config struct. If no file is found, the
// Save method is called to create a sample file
// before returning with an error.
func (c *Config) Load() error {
	err := jsonfile.Read(configFilePath, c)
	if err == nil || !os.IsNotExist(err) {
		return err
	}

	err = c.Save()
	if err != nil {
		return err
	}
	return errors.New("No 'gowrap.json' file found. " +
		"A sample file was created in the current directory. " +
		"Edit the file as required and re-run gowrap.")
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
