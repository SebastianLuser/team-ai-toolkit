package tokens

import (
	"net/http"
	"testing"

	"github.com/educabot/team-ai-toolkit/web"
)

func TestTenantInterceptor_SetsOrgID(t *testing.T) {
	interceptor := NewTenantInterceptor()

	req := web.NewMockRequest()
	SetClaims(req, &Claims{UserID: 1, OrgID: 5})
	nextCalled := false
	req.NextFn = func() { nextCalled = true }

	resp := interceptor(req)

	if resp.Status >= 400 {
		t.Errorf("should not return error, got status %d", resp.Status)
	}
	if !nextCalled {
		t.Error("Next() should have been called")
	}

	orgID, ok := req.Get("org_id")
	if !ok {
		t.Fatal("org_id should be set in context")
	}
	if orgID.(int64) != 5 {
		t.Errorf("org_id = %v, want 5", orgID)
	}
}

func TestTenantInterceptor_NoClaims(t *testing.T) {
	interceptor := NewTenantInterceptor()

	req := web.NewMockRequest()

	resp := interceptor(req)

	if resp.Status != http.StatusUnauthorized {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusUnauthorized)
	}
}

func TestTenantInterceptor_ZeroOrgID(t *testing.T) {
	interceptor := NewTenantInterceptor()

	req := web.NewMockRequest()
	SetClaims(req, &Claims{UserID: 1, OrgID: 0})

	resp := interceptor(req)

	if resp.Status != http.StatusForbidden {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusForbidden)
	}
}
