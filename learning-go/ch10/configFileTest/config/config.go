package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	APIKey string `json:"API_KEY"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	configFile, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		return config, err
	}

	return config, nil

}
