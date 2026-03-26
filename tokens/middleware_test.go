package tokens

import "testing"

func TestExtractBearerToken_Valid(t *testing.T) {
	got := extractBearerToken("Bearer abc123")
	if got != "abc123" {
		t.Errorf("extractBearerToken() = %q, want %q", got, "abc123")
	}
}

func TestExtractBearerToken_CaseInsensitive(t *testing.T) {
	got := extractBearerToken("bearer abc123")
	if got != "abc123" {
		t.Errorf("extractBearerToken() = %q, want %q", got, "abc123")
	}
}

func TestExtractBearerToken_Empty(t *testing.T) {
	got := extractBearerToken("")
	if got != "" {
		t.Errorf("extractBearerToken() = %q, want empty", got)
	}
}

func TestExtractBearerToken_NoBearerPrefix(t *testing.T) {
	got := extractBearerToken("Basic abc123")
	if got != "" {
		t.Errorf("extractBearerToken() = %q, want empty", got)
	}
}

func TestExtractBearerToken_OnlyBearer(t *testing.T) {
	got := extractBearerToken("Bearer")
	if got != "" {
		t.Errorf("extractBearerToken() = %q, want empty", got)
	}
}
