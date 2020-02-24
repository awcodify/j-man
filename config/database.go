package config

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/awcodify/j-man/utils"
)

// Database will store all database related config
type Database struct {
	Driver     string     `yaml:"driver"`
	PostgreSQL postgresql `yaml:"postgresql"`
}

type postgresql struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

// ConnectDB based on driver being used
func (cfg Config) ConnectDB() (db *sql.DB, err error) {
	if cfg.DB.Driver == "postgres" {
		options := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s",
			cfg.DB.PostgreSQL.Name,
			cfg.DB.PostgreSQL.User,
			cfg.DB.PostgreSQL.Password,
			cfg.DB.PostgreSQL.SSLMode,
		)
		db, err = sql.Open(cfg.DB.Driver, options)
		utils.DieIf(err)
	} else {
		err = errors.New("Driver not found")
	}

	return db, err
}
