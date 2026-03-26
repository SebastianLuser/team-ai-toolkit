package tokens

import (
	"net/http"
	"testing"

	"github.com/educabot/team-ai-toolkit/web"
)

func TestRequireRole_Allowed(t *testing.T) {
	interceptor := RequireRole("teacher", "admin")

	req := web.NewMockRequest()
	SetClaims(req, &Claims{UserID: 1, Roles: []string{"teacher"}})
	nextCalled := false
	req.NextFn = func() { nextCalled = true }

	resp := interceptor(req)

	if resp.Status >= 400 {
		t.Errorf("should not return error, got status %d", resp.Status)
	}
	if !nextCalled {
		t.Error("Next() should have been called")
	}
}

func TestRequireRole_Forbidden(t *testing.T) {
	interceptor := RequireRole("admin")

	req := web.NewMockRequest()
	SetClaims(req, &Claims{UserID: 1, Roles: []string{"teacher"}})

	resp := interceptor(req)

	if resp.Status != http.StatusForbidden {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusForbidden)
	}
}

func TestRequireRole_NoClaims(t *testing.T) {
	interceptor := RequireRole("teacher")

	req := web.NewMockRequest()

	resp := interceptor(req)

	if resp.Status != http.StatusUnauthorized {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusUnauthorized)
	}
}
