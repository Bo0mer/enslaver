package config

import (
	"os"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

type EnslaverConfig struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func LoadConfig(configFile string) (EnslaverConfig, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return EnslaverConfig{}, err
	}

	config := &EnslaverConfig{}

	decoder := candiedyaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return EnslaverConfig{}, err
	}

	return *config, nil
}
