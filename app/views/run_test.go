package views

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/awcodify/j-man/config"
	"github.com/stretchr/testify/assert"
)

func TestRunHandler(t *testing.T) {
	cfg, _ := config.New()
	cfg.HTML.Root = "html"
	cfg.HTML.Layout.BaseHTML = "html/layout/base.html"
	v := View{Config: cfg}

	rr := httptest.NewRecorder()
	req := &http.Request{Method: "GET"}

	v.RunHandler(rr, req)

	assert.HTTPBodyContains(t, v.RunHandler, "GET", "/run", nil, "Run Test!")
}
