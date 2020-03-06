package config

import (
	"github.com/go-redis/redis"
)

// Redis will store our cache
type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// ConnectRedis will initiate connection to redis
func (cfg Config) ConnectRedis() *redis.Client {
	options := &redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}

	return redis.NewClient(options)
}
