package config

import (
	"os"
	"testing"
)

func TestBaseConfig_IsProduction(t *testing.T) {
	tests := []struct {
		env  string
		want bool
	}{
		{"prod", true},
		{"staging", true},
		{"local", false},
		{"develop", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			cfg := BaseConfig{Env: tt.env}
			if got := cfg.IsProduction(); got != tt.want {
				t.Errorf("IsProduction(%q) = %v, want %v", tt.env, got, tt.want)
			}
		})
	}
}

func TestLoadBase_WithEnvVars(t *testing.T) {
	err := os.Setenv("PORT", "3000")
	if err != nil {
		return
	}
	err = os.Setenv("ENV", "staging")
	if err != nil {
		return
	}
	err = os.Setenv("DATABASE_URL", "postgres://localhost/test")
	if err != nil {
		return
	}
	err = os.Setenv("AUTH_PUBLIC_KEY", "test-key")
	if err != nil {
		return
	}
	err = os.Setenv("ALLOWED_ORIGINS", `https://a.com,https://b.com`)
	if err != nil {
		return
	}
	err = os.Setenv("API_KEY_BUGSNAG", "bugsnag-key")
	if err != nil {
		return
	}
	defer func() {
		err = os.Unsetenv("PORT")
		if err != nil {
			return
		}
		err = os.Unsetenv("ENV")
		if err != nil {
			return
		}
		err = os.Unsetenv("DATABASE_URL")
		if err != nil {
			return
		}
		err = os.Unsetenv("AUTH_PUBLIC_KEY")
		if err != nil {
			return
		}
		err = os.Unsetenv("ALLOWED_ORIGINS")
		if err != nil {
			return
		}
		err = os.Unsetenv("API_KEY_BUGSNAG")
		if err != nil {
			return
		}
	}()

	cfg := LoadBase()

	if cfg.Port != "3000" {
		t.Errorf("Port = %q, want %q", cfg.Port, "3000")
	}
	if cfg.Env != "staging" {
		t.Errorf("Env = %q, want %q", cfg.Env, "staging")
	}
	if cfg.DatabaseURL != "postgres://localhost/test" {
		t.Errorf("DatabaseURL = %q, want %q", cfg.DatabaseURL, "postgres://localhost/test")
	}
	if cfg.AuthPublicKey != "test-key" {
		t.Errorf("AuthPublicKey = %q, want %q", cfg.AuthPublicKey, "test-key")
	}
	if len(cfg.AllowedOrigins) != 2 {
		t.Errorf("AllowedOrigins length = %d, want 2", len(cfg.AllowedOrigins))
	}
	if cfg.BugsnagAPIKey != "bugsnag-key" {
		t.Errorf("BugsnagAPIKey = %q, want %q", cfg.BugsnagAPIKey, "bugsnag-key")
	}
}
