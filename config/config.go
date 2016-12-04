package config

import (
	"encoding/json"
	"os"
)

// Read decodes the configuration JSON file into the
// value pointed to by config.
func Read(filePath string, config interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(config)
	return err
}

// Save writes the value pointed to by config to a file.
func Save(filePath string, config interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	json, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}
	_, err = file.Write(json)
	return err
}

// ReadOrSave attempts to read the configuration file. If it
// cannot find the file, it will save a sample configuration
// file instead.
func ReadOrSave(filePath string, config interface{}) (readErr error, saveErr error) {
	readErr = Read(filePath, config)
	if readErr == nil {
		return readErr, nil
	}
	saveErr = Save(filePath, config)
	return readErr, saveErr
}
