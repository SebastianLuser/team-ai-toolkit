package web

import (
	"context"
	"fmt"
	"testing"
)

func TestNewMockRequest_Defaults(t *testing.T) {
	req := NewMockRequest()

	if req.Params == nil {
		t.Error("Params should be initialized")
	}
	if req.Queries == nil {
		t.Error("Queries should be initialized")
	}
	if req.Headers == nil {
		t.Error("Headers should be initialized")
	}
	if req.Values == nil {
		t.Error("Values should be initialized")
	}
	if req.Ctx == nil {
		t.Error("Ctx should be initialized")
	}
}

func TestMockRequest_Param(t *testing.T) {
	req := NewMockRequest()
	req.Params["id"] = "42"

	if got := req.Param("id"); got != "42" {
		t.Errorf("Param(id) = %q, want %q", got, "42")
	}
	if got := req.Param("missing"); got != "" {
		t.Errorf("Param(missing) = %q, want empty", got)
	}
}

func TestMockRequest_Query(t *testing.T) {
	req := NewMockRequest()
	req.Queries["page"] = "3"

	if got := req.Query("page"); got != "3" {
		t.Errorf("Query(page) = %q, want %q", got, "3")
	}
}

func TestMockRequest_Header(t *testing.T) {
	req := NewMockRequest()
	req.Headers["Authorization"] = "Bearer abc"

	if got := req.Header("Authorization"); got != "Bearer abc" {
		t.Errorf("Header(Authorization) = %q, want %q", got, "Bearer abc")
	}
}

func TestMockRequest_Context(t *testing.T) {
	req := NewMockRequest()
	if req.Context() != context.Background() {
		t.Error("Context() should return background context by default")
	}
}

func TestMockRequest_Bind_WithFn(t *testing.T) {
	req := NewMockRequest()
	req.BindFn = func(dest any) error {
		return fmt.Errorf("bind error")
	}

	err := req.Bind(nil)
	if err == nil || err.Error() != "bind error" {
		t.Errorf("Bind() error = %v, want 'bind error'", err)
	}
}

func TestMockRequest_Bind_NilFn(t *testing.T) {
	req := NewMockRequest()

	err := req.Bind(nil)
	if err != nil {
		t.Errorf("Bind() error = %v, want nil", err)
	}
}

func TestMockRequest_SetAndGet(t *testing.T) {
	req := NewMockRequest()
	req.Set("key", "value")

	v, ok := req.Get("key")
	if !ok {
		t.Error("Get(key) should return true")
	}
	if v != "value" {
		t.Errorf("Get(key) = %v, want %q", v, "value")
	}

	_, ok = req.Get("missing")
	if ok {
		t.Error("Get(missing) should return false")
	}
}

func TestMockRequest_Next_WithFn(t *testing.T) {
	req := NewMockRequest()
	called := false
	req.NextFn = func() { called = true }

	req.Next()
	if !called {
		t.Error("NextFn should have been called")
	}
}

func TestMockRequest_Next_NilFn(t *testing.T) {
	req := NewMockRequest()
	// Should not panic
	req.Next()
}
