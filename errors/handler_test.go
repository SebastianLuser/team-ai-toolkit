package errors

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandleError_MapsCorrectStatus(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
	}{
		{"not found", fmt.Errorf("user: %w", ErrNotFound), http.StatusNotFound},
		{"validation", fmt.Errorf("name: %w", ErrValidation), http.StatusBadRequest},
		{"unauthorized", ErrUnauthorized, http.StatusUnauthorized},
		{"forbidden", ErrForbidden, http.StatusForbidden},
		{"duplicate", fmt.Errorf("email: %w", ErrDuplicate), http.StatusConflict},
		{"conflict", ErrConflict, http.StatusConflict},
		{"unknown", fmt.Errorf("something unexpected"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := HandleError(tt.err)
			if resp.Status != tt.wantStatus {
				t.Errorf("HandleError(%v) status = %d, want %d", tt.err, resp.Status, tt.wantStatus)
			}
		})
	}
}
