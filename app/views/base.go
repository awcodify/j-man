package views

import (
	"context"
	"database/sql"

	"github.com/awcodify/j-man/config"
	"github.com/go-redis/redis"
)

// Config is all configurations managed by yaml
type View struct {
	*config.Config
	Ctx   context.Context
	Cache redis.Cmdable
	DB    *sql.DB
}

func (v View) getTemplatePath(name string) string {
	return v.HTML.Root + "/" + name + ".html"
}
