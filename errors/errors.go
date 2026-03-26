package errors

import "errors"

// Sentinel errors shared across all Educabot backend services.
// Projects can define additional domain-specific errors in their own providers/errors.go.

// ErrNotFound Not found
var (
	ErrNotFound = errors.New("not found")
)

// ErrValidation Validation
var (
	ErrValidation = errors.New("validation error")
)

// Auth
var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

// Conflict
var (
	ErrDuplicate = errors.New("duplicate entry")
	ErrConflict  = errors.New("conflict")
)

// Re-export stdlib errors functions so consumers don't need to import both packages.
var (
	Is  = errors.Is
	As  = errors.As
	New = errors.New
)
