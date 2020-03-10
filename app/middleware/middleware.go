package middleware

import (
	"github.com/go-redis/redis"
)

// Handler is HTTP handler
type Config struct {
	Cache redis.Cmdable
}
