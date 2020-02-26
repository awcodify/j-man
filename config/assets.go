package config

// HTML used for templating
type HTML struct {
	Root   string `yaml:"root"`
	Layout layout
}

type layout struct {
	Root     string `yaml:"root"`
	BaseHTML string `yaml:"base_html"`
}
