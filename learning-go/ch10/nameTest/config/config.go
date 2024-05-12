package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	LogFilePath string `json:"logFilePath"`
}

func LoadConfiguration(file string) (Configuration, error) {
	var config Configuration
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
