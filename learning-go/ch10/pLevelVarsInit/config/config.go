package config

import (
	"log"
	"os"
)

var configSettings map[string]string

func init() {
	configSettings = loadConfigSettings()
}

func loadConfigSettings() map[string]string {
	settings := make(map[string]string)

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalf("API_KEY environment variable is not set.")
	}
	settings["apiKey"] = apiKey

	return settings
}

// GetConfig provides a safe way to access configuration settings.
func GetConfig(key string) string {
	value, exists := configSettings[key]
	if !exists {
		log.Printf("config key %s not found", key)
	}

	return value
}
