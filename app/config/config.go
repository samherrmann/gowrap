package config

import (
	"errors"

	"os"

	"github.com/samherrmann/gowrap/jsonfile"
)

const (
	configFilePath = "gowrap.json"
)

// Config represents the application configuration.
type Config struct {
	Targets *[]string `json:"targets"`
}

// New returns a Config struct.
func New() *Config {
	return new(Config)
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
		return nil
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
