package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/awcodify/j-man/app/services/test"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func ServeAPI() {
	cfg, err := config.New()
	utils.DieIf(err)

	db, err := cfg.ConnectDB()
	utils.DieIf(err)

	wr := RequestWrapper{
		config: cfg,
		db:     db,
		ctx:    context.Background(),
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/rounds", func(r chi.Router) {
		r.Post("/", wr.CreateRounds)
	})
	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.App.Server.Host, cfg.App.Server.Port), r)
}

func (rn *RoundRequest) Bind(r *http.Request) error {
	// a.Article is nil if no Article fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if rn.Round == nil {
		return errors.New("missing required round fields.")
	}

	rn.ProtectedID = "" // unset the protected ID
	return nil
}

func (wr RequestWrapper) CreateRounds(w http.ResponseWriter, r *http.Request) {
	data := &RoundRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	round := data.Round
	runner := test.Runner{
		Config: wr.config,
		Round:  round,
		DB:     wr.db,
	}

	runner.Run(wr.ctx)
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
