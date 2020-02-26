package views

import (
	"github.com/awcodify/j-man/config"
)

var cfg Config

// Config is all configurations managed by yaml
type Config struct {
	HTML config.HTML
}

func init() {
	c := config.New()
	cfg.HTML = c.HTML
}

func getTemplatePath(name string) string {
	return cfg.HTML.Root + "/" + name + ".html"
}
