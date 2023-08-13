package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// errors
var ErrPathNotProvided error = errors.New("config path not provided")
var ErrUnableToRead error = errors.New("unable to read config file")
var ErrUnableToParse error = errors.New("unable to parse config file yaml")

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func Read(path string) (*Config, error) {
	if path == "" {
		return nil, ErrPathNotProvided
	}

	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, ErrUnableToRead
	}

	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, ErrUnableToParse
	}

	return &config, nil
}
