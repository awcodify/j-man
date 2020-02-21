package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config will be seperated between production and development
type Config struct {
	// Server is application server
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	HTML HTML
}

// New will instantiatea config from production or development
func New() (c Config, err error) {
	var cfg Config
	var file *os.File
	env := os.Getenv("APP_ENV")

	if env != "" {
		fileName := fmt.Sprintf("config.%s.yaml", strings.ToLower(env))
		file, err = os.Open(fileName)
	} else {
		return Config{}, errors.New("Please choose production or development")
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
