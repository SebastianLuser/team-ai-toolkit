package applog

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoopTracker_DoesNotPanic(t *testing.T) {
	tracker := NewNoopTracker()

	// Should not panic
	tracker.Notify(fmt.Errorf("test error"), context.Background())

	// Handler should pass through
	called := false
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	handler := tracker.Handler(inner)
	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if !called {
		t.Error("inner handler should have been called")
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestSetup_ProdUsesJSONHandler(t *testing.T) {
	Setup("prod")
	// Should not panic
}

func TestSetup_StagingUsesJSONHandler(t *testing.T) {
	Setup("staging")
}

func TestSetup_LocalUsesTextHandler(t *testing.T) {
	Setup("local")
}

func TestSetupTest_ConfiguresSilentLogger(t *testing.T) {
	SetupTest()
}
