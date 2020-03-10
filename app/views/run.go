package views

import (
	"html/template"
	"net/http"

	"github.com/awcodify/j-man/utils"
)

// PageData will be used by html page
type PageData struct {
	Title string
}

// RunHandler will handle http request for `/run`
func (cfg Config) RunHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		IndexRun(w, cfg)
	}
}

// IndexRun is index page of Run
func IndexRun(w http.ResponseWriter, cfg Config) {
	t, err := template.New("").ParseFiles(cfg.getTemplatePath("run"), cfg.HTML.Layout.BaseHTML)
	utils.DieIf(err)

	pageData := PageData{Title: "Run Test!"}
	err = t.ExecuteTemplate(w, "base", pageData)
	utils.DieIf(err)
}
