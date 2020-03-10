package views

import (
	"github.com/awcodify/j-man/config"
)

// Config is all configurations managed by yaml
type View struct {
	config.Config
}

func (v View) getTemplatePath(name string) string {
	return v.HTML.Root + "/" + name + ".html"
}
