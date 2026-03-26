package errors

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/educabot/team-ai-toolkit/web"
)

// HandleError maps a domain error to an HTTP response.
// Checks the error chain with errors.Is() and returns the appropriate status + code.
// Projects can wrap this to add their own domain-specific error mappings.
func HandleError(err error) web.Response {
	switch {
	case errors.Is(err, ErrNotFound):
		return web.Err(http.StatusNotFound, "not_found", err.Error())

	case errors.Is(err, ErrValidation):
		return web.Err(http.StatusBadRequest, "validation_error", err.Error())

	case errors.Is(err, ErrUnauthorized):
		return web.Err(http.StatusUnauthorized, "unauthorized", err.Error())

	case errors.Is(err, ErrForbidden):
		return web.Err(http.StatusForbidden, "forbidden", err.Error())

	case errors.Is(err, ErrDuplicate):
		return web.Err(http.StatusConflict, "duplicate", err.Error())

	case errors.Is(err, ErrConflict):
		return web.Err(http.StatusConflict, "conflict", err.Error())

	default:
		slog.Error("unhandled error", "err", err)
		return web.Err(http.StatusInternalServerError, "internal_error", "something went wrong")
	}
}
