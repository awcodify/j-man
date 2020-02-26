package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/awcodify/j-man/app/views"
	"github.com/awcodify/j-man/config"
)

func main() {
	cfg := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/run", views.RunHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), logRequest(mux))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
