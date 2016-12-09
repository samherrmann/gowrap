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
		Targets:   &[]string{t.String()},
		platforms: &[]gotools.Platform{},
	}
	return c
}

// Config represents the application configuration.
type Config struct {
	Targets *[]string `json:"targets"`

	platforms *[]gotools.Platform
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
	if err == nil {
		return c.validate()
	}
	if !os.IsNotExist(err) {
		return err
	}

	err = c.Save()
	if err != nil {
		return err
	}
	return errors.New("No '" + configFilePath + "' file found. " +
		"A sample file was created in the current directory. " +
		"Edit the file as required and re-run gowrap.")
}

// validate verifies that the target values
// are of proper format by unmarshaling them
// to gotools/Plaform types.
func (c *Config) validate() error {
	ps := &[]gotools.Platform{}

	for _, t := range *c.Targets {
		p := gotools.NewPlatform()
		if err := p.Unmarshal(t); err != nil {
			return err
		}
		*ps = append(*ps, *p)
	}
	c.platforms = ps
	return nil
}

// Platforms returns the targets as a
// gotools/Platform slice.
func (c *Config) Platforms() *[]gotools.Platform {
	return c.platforms
}
