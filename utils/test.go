package utils

import (
	"github.com/khaiql/dbcleaner"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

func CleanTablesAfter(dsn string, models []string, fn func()) {
	cleaner := dbcleaner.New()
	psql := engine.NewPostgresEngine(dsn)
	cleaner.SetEngine(psql)

	for _, model := range models {
		cleaner.Acquire(model)
	}

	fn()

	for _, model := range models {
		cleaner.Clean(model)
	}
}
