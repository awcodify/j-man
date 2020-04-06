package app

import (
	"context"
	"database/sql"

	"github.com/awcodify/j-man/app/models"
	"github.com/awcodify/j-man/config"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

type RequestWrapper struct {
	config *config.Config
	db     *sql.DB
	ctx    context.Context
}

type RoundRequest struct {
	*models.Round
	ProtectedID string `json:"id"`
}

type RoundResponse struct {
	*models.Round
	Elapsed int64 `json:"elapsed"`
}
