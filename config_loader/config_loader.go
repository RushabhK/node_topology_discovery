package config_loader

import (
	"encoding/json"
	"os"
)

type ConfigLoader interface {
	Load(filePath string) (ConfigData, error)
}

type configLoader struct {
}

func NewConfigLoader() ConfigLoader {
	return configLoader{}
}

func (configLoader configLoader) Load(filePath string) (ConfigData, error) {
	configuration := ConfigData{}
	file, err := os.Open(filePath)
	if err != nil {
		return configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	return configuration, err
}
