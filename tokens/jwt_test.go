package tokens

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateTestKeyPair(t *testing.T) (*rsa.PrivateKey, *rsa.PublicKey) {
	t.Helper()
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}
	return privateKey, &privateKey.PublicKey
}

func signTestToken(t *testing.T, privateKey *rsa.PrivateKey, claims *Claims) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(privateKey)
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}
	return signed
}

func TestValidateJWT_ValidToken(t *testing.T) {
	privKey, pubKey := generateTestKeyPair(t)

	original := &Claims{
		UserID: 42,
		OrgID:  1,
		Roles:  []string{"teacher", "coordinator"},
		Email:  "carlos@school.edu",
		Name:   "Carlos",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenStr := signTestToken(t, privKey, original)

	got, err := ValidateJWT(tokenStr, pubKey)
	if err != nil {
		t.Fatalf("ValidateJWT() error = %v", err)
	}

	if got.UserID != 42 {
		t.Errorf("UserID = %d, want 42", got.UserID)
	}
	if got.OrgID != 1 {
		t.Errorf("OrgID = %d, want 1", got.OrgID)
	}
	if got.Email != "carlos@school.edu" {
		t.Errorf("Email = %q, want %q", got.Email, "carlos@school.edu")
	}
	if len(got.Roles) != 2 {
		t.Errorf("Roles length = %d, want 2", len(got.Roles))
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	privKey, pubKey := generateTestKeyPair(t)

	claims := &Claims{
		UserID: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}

	tokenStr := signTestToken(t, privKey, claims)

	_, err := ValidateJWT(tokenStr, pubKey)
	if err == nil {
		t.Error("ValidateJWT() should fail on expired token")
	}
}

func TestValidateJWT_WrongKey(t *testing.T) {
	privKey, _ := generateTestKeyPair(t)
	_, wrongPubKey := generateTestKeyPair(t)

	claims := &Claims{
		UserID: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	tokenStr := signTestToken(t, privKey, claims)

	_, err := ValidateJWT(tokenStr, wrongPubKey)
	if err == nil {
		t.Error("ValidateJWT() should fail with wrong public key")
	}
}

func TestValidateJWT_InvalidTokenString(t *testing.T) {
	_, pubKey := generateTestKeyPair(t)

	_, err := ValidateJWT("not.a.valid.token", pubKey)
	if err == nil {
		t.Error("ValidateJWT() should fail on invalid token string")
	}
}

func TestParseRSAPublicKey_Valid(t *testing.T) {
	privKey, _ := generateTestKeyPair(t)
	pubBytes := x509.MarshalPKCS1PublicKey(&privKey.PublicKey)
	pkixBytes, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		t.Fatalf("failed to marshal public key: %v", err)
	}
	_ = pubBytes

	pemBlock := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pkixBytes,
	})

	key, err := ParseRSAPublicKey(string(pemBlock))
	if err != nil {
		t.Fatalf("ParseRSAPublicKey() error = %v", err)
	}
	if key == nil {
		t.Error("ParseRSAPublicKey() returned nil key")
	}
}

func TestParseRSAPublicKey_InvalidPEM(t *testing.T) {
	_, err := ParseRSAPublicKey("not a pem")
	if err == nil {
		t.Error("ParseRSAPublicKey() should fail on invalid PEM")
	}
}

func TestParseRSAPublicKey_InvalidKey(t *testing.T) {
	pemBlock := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: []byte("not a real key"),
	})

	_, err := ParseRSAPublicKey(string(pemBlock))
	if err == nil {
		t.Error("ParseRSAPublicKey() should fail on invalid key bytes")
	}
}
