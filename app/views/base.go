package views

import (
	"github.com/awcodify/j-man/config"
)

// Config is all configurations managed by yaml
type Config struct {
	config.Config
}

func (cfg Config) getTemplatePath(name string) string {
	return cfg.HTML.Root + "/" + name + ".html"
}
