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
func (v View) RunHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		IndexRun(w, v)
	}
}

// IndexRun is index page of Run
func IndexRun(w http.ResponseWriter, v View) {
	t, err := template.New("").ParseFiles(v.getTemplatePath("run"), v.HTML.Layout.BaseHTML)
	utils.DieIf(err)

	pageData := PageData{Title: "Run Test!"}
	err = t.ExecuteTemplate(w, "base", pageData)
	utils.DieIf(err)
}
