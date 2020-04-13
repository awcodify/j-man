package views

import (
	"log"
	"testing"

	"github.com/awcodify/j-man/config"
	"github.com/stretchr/testify/assert"
)

func Test_getTemplatePath(t *testing.T) {
	cfg, err := config.New()
	log.Println(err)
	cfg.HTML.Root = "app/views/html"
	v := View{Config: cfg}

	actual := v.getTemplatePath("run")

	assert.Equal(t, "app/views/html/run.html", actual)
}
