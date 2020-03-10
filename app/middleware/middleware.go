package middleware

import (
	"github.com/go-redis/redis"
)

// Handler is HTTP handler
type Middleware struct {
	Cache redis.Cmdable
}
