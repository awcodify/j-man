package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/awcodify/j-man/app/middleware"
	"github.com/awcodify/j-man/app/views"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/utils"
)

func Run() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	cache := cfg.ConnectRedis()

	pong, err := cache.Ping().Result()
	utils.DieIf(err)
	log.Println(pong)

	db, err := cfg.ConnectDB()
	utils.DieIf(err)

	v := views.View{Config: cfg,
		Ctx:   context.Background(),
		DB:    db,
		Cache: cache,
	}
	midd := middleware.Middleware{Cache: cache}

	mux := http.NewServeMux()
	mux.HandleFunc("/sign_in", v.HandleSignIn)
	mux.HandleFunc("/authenticate", v.Authenticate)
	mux.HandleFunc("/run", midd.Auth(v.RunHandler))

	return http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), logRequest(mux))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
