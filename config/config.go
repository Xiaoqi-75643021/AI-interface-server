package config

import (
	"encoding/json"
	"os"
)

// Configuration holds all the configuration for both OpenAI and Baidu services
type Configuration struct {
	OpenAI struct {
		APIKey      string            `json:"APIKey"`
		APIURL      string            `json:"APIURL"`
		LogFileName string            `json:"LogFileName"`
		ModelsList  map[string]bool   `json:"ModelsList"`
	} `json:"OpenAI"`
	Baidu struct {
		APIKey    string `json:"APIKey"`
		SecretKey string `json:"SecretKey"`
		BaseURL   string `json:"BaseURL"`
	} `json:"Baidu"`
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
