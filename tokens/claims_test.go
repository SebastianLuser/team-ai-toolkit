package tokens

import "testing"

func TestClaims_HasRole(t *testing.T) {
	claims := Claims{Roles: []string{"teacher", "coordinator"}}

	if !claims.HasRole("teacher") {
		t.Error("should have role teacher")
	}
	if !claims.HasRole("coordinator") {
		t.Error("should have role coordinator")
	}
	if claims.HasRole("admin") {
		t.Error("should not have role admin")
	}
}

func TestClaims_HasAnyRole(t *testing.T) {
	claims := Claims{Roles: []string{"teacher"}}

	if !claims.HasAnyRole("admin", "teacher") {
		t.Error("should match at least one role")
	}
	if claims.HasAnyRole("admin", "coordinator") {
		t.Error("should not match any role")
	}
}

func TestClaims_HasRole_EmptyRoles(t *testing.T) {
	claims := Claims{Roles: []string{}}

	if claims.HasRole("teacher") {
		t.Error("empty roles should not match anything")
	}
	if claims.HasAnyRole("teacher", "admin") {
		t.Error("empty roles should not match any")
	}
}
