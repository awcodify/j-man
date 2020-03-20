package utils

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCleanTablesAfter(t *testing.T) {
	dsn := "user=jmanager_user password=password dbname=jmanager sslmode=disable"
	db, _ := sql.Open("postgres", dsn)

	CleanTablesAfter(dsn, []string{"users"}, func() {
		db.Exec(`INSERT INTO users (email) VALUES ("test@example.com")`)
	})

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users")
	row.Scan(&count)

	assert.Equal(t, 0, count)
}
