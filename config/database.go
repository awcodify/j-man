package config

import (
	"database/sql"

	"github.com/awcodify/j-man/utils"
	// It used for postgresql driver
	_ "github.com/lib/pq"
)

// Database will used for managing our state and scripts
type Database struct {
	DSN string `yaml:"dsn"`
}

// ConnectDB based on driver being used
func (cfg Config) ConnectDB() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", cfg.DB.DSN)
	utils.DieIf(err)

	return db, err
}
