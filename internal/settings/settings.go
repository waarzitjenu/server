package settings

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Port  uint `json:"port"`
	Debug bool `json:"debug"`
}

var InvalidPort = errors.New("invalid port")

func Read(filename string) (*Config, error) {
	settingsFile, err := os.Open(filename)

	if err != nil {
		homeDir, err := os.UserHomeDir()
		settingsFile, err = os.Open(homeDir + "/" + filename)

		if err != nil {
			return nil, err
		}
	}

	byteValue, err := ioutil.ReadAll(settingsFile)

	if err != nil {
		return nil, err
	}

	err = settingsFile.Close()

	if err != nil {
		return nil, err
	}

	var configFile Config

	err = json.Unmarshal(byteValue, &configFile)

	if err != nil {
		return nil, err
	}

	// Validate entered port
	if configFile.Port == 0 {
		return nil, InvalidPort
	}

	if configFile.Debug {
		log.Printf("Successfully parsed settings file: %s \n", filename)
	}

	return &configFile, nil
}

func Write(filename string, config *Config) error {
	configBytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, configBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func IsCorrupted(filename string) bool {
	_, err := Read(filename)

	if err == InvalidPort {
		return true
	}

	return false
}
