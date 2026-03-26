package tokens

import (
	"testing"

	"github.com/educabot/team-ai-toolkit/web"
)

func TestSetAndGetClaims(t *testing.T) {
	req := web.NewMockRequest()
	claims := &Claims{UserID: 42, OrgID: 1, Email: "test@test.com"}

	SetClaims(req, claims)

	got := GetClaims(req)
	if got == nil {
		t.Fatal("GetClaims() returned nil")
	}
	if got.UserID != 42 {
		t.Errorf("UserID = %d, want 42", got.UserID)
	}
	if got.OrgID != 1 {
		t.Errorf("OrgID = %d, want 1", got.OrgID)
	}
}

func TestGetClaims_ReturnsNilWhenMissing(t *testing.T) {
	req := web.NewMockRequest()

	got := GetClaims(req)
	if got != nil {
		t.Errorf("GetClaims() = %v, want nil", got)
	}
}

func TestMustClaims_PanicsWhenMissing(t *testing.T) {
	req := web.NewMockRequest()

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustClaims() should panic when claims are missing")
		}
	}()

	MustClaims(req)
}

func TestMustClaims_ReturnsClaims(t *testing.T) {
	req := web.NewMockRequest()
	claims := &Claims{UserID: 99, OrgID: 5}
	SetClaims(req, claims)

	got := MustClaims(req)
	if got.UserID != 99 {
		t.Errorf("UserID = %d, want 99", got.UserID)
	}
}

func TestUserID_ExtractsFromClaims(t *testing.T) {
	req := web.NewMockRequest()
	SetClaims(req, &Claims{UserID: 77, OrgID: 3})

	if got := UserID(req); got != 77 {
		t.Errorf("UserID() = %d, want 77", got)
	}
}

func TestOrgID_ExtractsFromClaims(t *testing.T) {
	req := web.NewMockRequest()
	SetClaims(req, &Claims{UserID: 1, OrgID: 88})

	if got := OrgID(req); got != 88 {
		t.Errorf("OrgID() = %d, want 88", got)
	}
}

func TestGetClaims_ReturnsNilOnWrongType(t *testing.T) {
	req := web.NewMockRequest()
	req.Set("auth_claims", "not-a-claims-pointer")

	got := GetClaims(req)
	if got != nil {
		t.Errorf("GetClaims() = %v, want nil for wrong type", got)
	}
}
