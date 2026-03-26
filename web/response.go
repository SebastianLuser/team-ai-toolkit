package web

import "net/http"

// Response represents an HTTP response returned by handlers.
type Response struct {
	Status int
	Body   any
}

// JSON creates a success response with a status code and body.
func JSON(status int, body any) Response {
	return Response{Status: status, Body: body}
}

// OK creates a 200 response wrapping data in {"data": ...}.
func OK(data any) Response {
	return JSON(http.StatusOK, map[string]any{"data": data})
}

// Created creates a 201 response wrapping data in {"data": ...}.
func Created(data any) Response {
	return JSON(http.StatusCreated, map[string]any{"data": data})
}

// NoContent creates a 204 response without body.
func NoContent() Response {
	return Response{Status: http.StatusNoContent}
}

// Err creates an error response with a code and message.
func Err(status int, code, message string) Response {
	return Response{
		Status: status,
		Body: map[string]any{
			"error": map[string]string{
				"code":    code,
				"message": message,
			},
		},
	}
}

// Paginated creates a paginated response.
func Paginated(data any, total, page, perPage int) Response {
	return JSON(http.StatusOK, map[string]any{
		"data":     data,
		"total":    total,
		"page":     page,
		"per_page": perPage,
	})
}
