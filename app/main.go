package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/awcodify/j-man/app/middleware"
	"github.com/awcodify/j-man/app/views"
	"github.com/awcodify/j-man/config"
)

func main() {
	cfg := config.New()
	cache := cfg.ConnectRedis()
	v := views.Config{Config: cfg}
	midd := middleware.Config{Cache: cache}

	mux := http.NewServeMux()
	mux.HandleFunc("/sign_in", v.HandleSignIn)
	mux.HandleFunc("/authenticate", v.Authenticate)
	mux.HandleFunc("/run", midd.Auth(v.RunHandler))
	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.App.Server.Host, cfg.App.Server.Port), logRequest(mux))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
