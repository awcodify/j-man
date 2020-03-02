package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/awcodify/j-man/utils"
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
			Path string `yaml:"host"`
		}
		Server struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}
	}
	HTML HTML     `yaml:"HTML"`
	DB   Database `yaml:"DB"`
}

// New will instantiate config from production or development
func New() (c Config) {
	env := os.Getenv("APP_ENV")

	fileName := fmt.Sprintf("config.%s.yaml", strings.ToLower(env))
	file, err := os.Open(basepath + "/" + fileName)
	utils.DieIf(err)

	if env == "" {
		panic(errors.New("Please choose production or development"))
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&c)
	utils.DieIf(err)

	return c
}
