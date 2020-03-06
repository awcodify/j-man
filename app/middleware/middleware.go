package middleware

import (
	"github.com/awcodify/j-man/config"
	"github.com/go-redis/redis"
)

// Handler is HTTP handler
type Handler struct {
	Config config.Config
	Cache  redis.Cmdable
}
