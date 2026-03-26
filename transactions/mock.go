package transactions

import (
	"context"
	"database/sql"
)

// MockDBTX is a test double for DBTX. Override the functions you need per test.
type MockDBTX struct {
	GetContextFn       func(ctx context.Context, dest any, query string, args ...any) error
	SelectContextFn    func(ctx context.Context, dest any, query string, args ...any) error
	ExecContextFn      func(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContextFn  func(ctx context.Context, query string, args ...any) *sql.Row
}

func (m *MockDBTX) GetContext(ctx context.Context, dest any, query string, args ...any) error {
	if m.GetContextFn != nil {
		return m.GetContextFn(ctx, dest, query, args...)
	}
	return nil
}

func (m *MockDBTX) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	if m.SelectContextFn != nil {
		return m.SelectContextFn(ctx, dest, query, args...)
	}
	return nil
}

func (m *MockDBTX) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if m.ExecContextFn != nil {
		return m.ExecContextFn(ctx, query, args...)
	}
	return nil, nil
}

func (m *MockDBTX) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if m.QueryRowContextFn != nil {
		return m.QueryRowContextFn(ctx, query, args...)
	}
	return nil
}
