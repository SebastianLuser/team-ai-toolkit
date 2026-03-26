package dbconn

import (
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// MustConnect creates a new PostgreSQL connection pool or panics on failure.
// The pool is configured with sensible defaults for a web service.
func MustConnect(dsn string) *sqlx.DB {
	db := sqlx.MustConnect("postgres", dsn)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	slog.Info("database connected", "max_open", 25, "max_idle", 10)
	return db
}
