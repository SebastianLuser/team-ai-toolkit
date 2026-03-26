package web

import "context"

// Request is a framework-agnostic HTTP request interface.
// Handlers and middleware depend on this, never on a specific framework.
type Request interface {
	// Param returns a path parameter value (e.g. :id).
	Param(key string) string

	// Query returns a query string parameter value (e.g. ?page=1).
	Query(key string) string

	// Header returns a request header value.
	Header(key string) string

	// Bind decodes the JSON body into dest.
	Bind(dest any) error

	// Set stores a value in the request context (e.g. claims).
	Set(key string, value any)

	// Get retrieves a value from the request context.
	Get(key string) (any, bool)

	// Context returns the request context for cancellation and deadlines.
	Context() context.Context

	// Next calls the next handler/middleware in the chain.
	Next()
}
