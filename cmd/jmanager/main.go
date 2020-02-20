package main

import (
	"fmt"
	"net/http"

	"github.com/awcodify/j-man/app/views"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/utils"
)

func main() {
	cfg, err := config.New()
	utils.DieIf(err)

	app := http.NewServeMux()

	v := views.Config{HTML: cfg.HTML}
	app.HandleFunc("/run", v.RunHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), app)
}
