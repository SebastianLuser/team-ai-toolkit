package applog

import (
	"context"
	"net/http"
)

// ErrorTracker is an abstract interface for error tracking services.
// Today it's implemented by Bugsnag. Tomorrow it could be Sentry, Datadog, etc.
// The pattern is the same as web/gin: abstract interface + concrete adapter.
type ErrorTracker interface {
	// Notify reports an error to the tracking service with optional context.
	Notify(err error, ctx context.Context)

	// Handler returns an http.Handler middleware for automatic error capture.
	// This integrates with the HTTP framework at the net/http level.
	Handler(next http.Handler) http.Handler
}

// NoopTracker is a no-op implementation for development/testing.
type NoopTracker struct{}

func NewNoopTracker() *NoopTracker                                     { return &NoopTracker{} }
func (t *NoopTracker) Notify(_ error, _ context.Context)               {}
func (t *NoopTracker) Handler(next http.Handler) http.Handler          { return next }
