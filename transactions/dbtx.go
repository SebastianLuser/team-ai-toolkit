package transactions

import (
	"context"
	"database/sql"
)

// DBTX is an interface shared by *sqlx.DB and *sqlx.Tx.
// Repositories accept this so they can work inside or outside a transaction.
type DBTX interface {
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
