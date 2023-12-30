package config

import (
	"encoding/json"
	"os"
)

type APIConfig struct {
	Name        string            		`json:"Name"`
	Categories  map[string][]Service 	`json:"Categories"`
}

type Service struct {
	InterfaceName string            `json:"InterfaceName"`
	APIKey        string            `json:"APIKey"`
	SecretKey     string            `json:"SecretKey,omitempty"`
	APIURL        string            `json:"APIURL"`
	ModelsList    map[string]bool   `json:"ModelsList,omitempty"`
}

type Configuration struct {
	APIs []APIConfig `json:"APIs"`
}

// LoadConfig reads a config file from a given path and unmarshals it into a Configuration struct
func LoadConfig(path string) (*Configuration, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &Configuration{}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
