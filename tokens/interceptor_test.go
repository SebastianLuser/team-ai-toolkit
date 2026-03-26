package tokens

import (
	"crypto/rsa"
	"net/http"
	"testing"
	"time"

	"github.com/educabot/team-ai-toolkit/web"
	"github.com/golang-jwt/jwt/v5"
)

func setupAuthTest(t *testing.T) (*rsa.PrivateKey, *rsa.PublicKey, string) {
	t.Helper()
	privKey, pubKey := generateTestKeyPair(t)

	claims := &Claims{
		UserID: 42,
		OrgID:  1,
		Roles:  []string{"teacher"},
		Email:  "test@school.edu",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenStr := signTestToken(t, privKey, claims)
	return privKey, pubKey, tokenStr
}

func TestAuthInterceptor_ValidToken(t *testing.T) {
	_, pubKey, tokenStr := setupAuthTest(t)

	interceptor := NewAuthInterceptor(pubKey)

	req := web.NewMockRequest()
	req.Headers["Authorization"] = "Bearer " + tokenStr
	nextCalled := false
	req.NextFn = func() { nextCalled = true }

	resp := interceptor(req)

	if resp.Status >= 400 {
		t.Errorf("should not return error, got status %d", resp.Status)
	}
	if !nextCalled {
		t.Error("Next() should have been called")
	}

	claims := GetClaims(req)
	if claims == nil {
		t.Fatal("claims should be set in context")
	}
	if claims.UserID != 42 {
		t.Errorf("UserID = %d, want 42", claims.UserID)
	}
}

func TestAuthInterceptor_MissingToken(t *testing.T) {
	_, pubKey := generateTestKeyPair(t)
	interceptor := NewAuthInterceptor(pubKey)

	req := web.NewMockRequest()

	resp := interceptor(req)

	if resp.Status != http.StatusUnauthorized {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusUnauthorized)
	}
}

func TestAuthInterceptor_InvalidToken(t *testing.T) {
	_, pubKey := generateTestKeyPair(t)
	interceptor := NewAuthInterceptor(pubKey)

	req := web.NewMockRequest()
	req.Headers["Authorization"] = "Bearer invalid.token.here"

	resp := interceptor(req)

	if resp.Status != http.StatusUnauthorized {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusUnauthorized)
	}
}

func TestAuthInterceptor_ExpiredToken(t *testing.T) {
	privKey, pubKey := generateTestKeyPair(t)

	claims := &Claims{
		UserID: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
		},
	}
	tokenStr := signTestToken(t, privKey, claims)

	interceptor := NewAuthInterceptor(pubKey)
	req := web.NewMockRequest()
	req.Headers["Authorization"] = "Bearer " + tokenStr

	resp := interceptor(req)

	if resp.Status != http.StatusUnauthorized {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusUnauthorized)
	}
}

func TestOptionalAuthInterceptor_NoToken(t *testing.T) {
	_, pubKey := generateTestKeyPair(t)
	interceptor := NewOptionalAuthInterceptor(pubKey)

	req := web.NewMockRequest()
	nextCalled := false
	req.NextFn = func() { nextCalled = true }

	resp := interceptor(req)

	if resp.Status >= 400 {
		t.Errorf("should not return error, got status %d", resp.Status)
	}
	if !nextCalled {
		t.Error("Next() should have been called even without token")
	}
	if GetClaims(req) != nil {
		t.Error("claims should be nil when no token")
	}
}

func TestOptionalAuthInterceptor_ValidToken(t *testing.T) {
	_, pubKey, tokenStr := setupAuthTest(t)
	interceptor := NewOptionalAuthInterceptor(pubKey)

	req := web.NewMockRequest()
	req.Headers["Authorization"] = "Bearer " + tokenStr
	nextCalled := false
	req.NextFn = func() { nextCalled = true }

	interceptor(req)

	if !nextCalled {
		t.Error("Next() should have been called")
	}

	claims := GetClaims(req)
	if claims == nil {
		t.Fatal("claims should be set with valid token")
	}
	if claims.UserID != 42 {
		t.Errorf("UserID = %d, want 42", claims.UserID)
	}
}

func TestOptionalAuthInterceptor_InvalidToken(t *testing.T) {
	_, pubKey := generateTestKeyPair(t)
	interceptor := NewOptionalAuthInterceptor(pubKey)

	req := web.NewMockRequest()
	req.Headers["Authorization"] = "Bearer bad.token"
	nextCalled := false
	req.NextFn = func() { nextCalled = true }

	interceptor(req)

	if !nextCalled {
		t.Error("Next() should still be called with invalid token (optional)")
	}
	if GetClaims(req) != nil {
		t.Error("claims should be nil with invalid token")
	}
}
