package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Join(filepath.Dir(b), "../")
)

// Config will be seperated between production and development
type Config struct {
	App struct {
		JMeter struct {
			Path string `yaml:"path"`
		}
		Server struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}
	}
	HTML  HTML     `yaml:"HTML"`
	DB    Database `yaml:"DB"`
	OAuth OAuth    `yaml:"OAuth"`
	Redis Redis    `yaml:"Redis"`
}

// New will instantiate config from production or development
func New() (c *Config, err error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return nil, errors.New("Please choose production or development")
	}

	fileName := fmt.Sprintf("config.%s.yaml", strings.ToLower(env))
	file, err := os.Open(basepath + "/" + fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
