package config

import (
	"os"
	"testing"
)

func TestEnvOr_ReturnsValue(t *testing.T) {
	err := os.Setenv("TEST_ENV_OR", "hello")
	if err != nil {
		return
	}
	defer func() {
		err := os.Unsetenv("TEST_ENV_OR")
		if err != nil {

		}
	}()

	got := EnvOr("TEST_ENV_OR", "fallback")
	if got != "hello" {
		t.Errorf("EnvOr() = %q, want %q", got, "hello")
	}
}

func TestEnvOr_ReturnsFallback(t *testing.T) {
	err := os.Unsetenv("TEST_ENV_OR_MISSING")
	if err != nil {
		return
	}

	got := EnvOr("TEST_ENV_OR_MISSING", "fallback")
	if got != "fallback" {
		t.Errorf("EnvOr() = %q, want %q", got, "fallback")
	}
}

func TestMustEnv_ReturnsValue(t *testing.T) {
	err := os.Setenv("TEST_MUST_ENV", "secret")
	if err != nil {
		return
	}
	defer func() {
		err := os.Unsetenv("TEST_MUST_ENV")
		if err != nil {
		}
	}()

	got := MustEnv("TEST_MUST_ENV")
	if got != "secret" {
		t.Errorf("MustEnv() = %q, want %q", got, "secret")
	}
}

func TestMustEnv_PanicsWhenMissing(t *testing.T) {
	err := os.Unsetenv("TEST_MUST_ENV_MISSING")
	if err != nil {
		return

	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("MustEnv() did not panic on missing var")
		}
	}()

	MustEnv("TEST_MUST_ENV_MISSING")
}

func TestEnvSplit_ReturnsSlice(t *testing.T) {
	err := os.Setenv("TEST_ENV_SPLIT", "a,b,c")
	if err != nil {
		return

	}
	defer func() {
		err := os.Unsetenv("TEST_ENV_SPLIT")
		if err != nil {

		}
	}()

	got := EnvSplit("TEST_ENV_SPLIT", ",", nil)
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("EnvSplit() = %v, want [a b c]", got)
	}
}

func TestEnvSplit_ReturnsFallback(t *testing.T) {
	err := os.Unsetenv("TEST_ENV_SPLIT_MISSING")
	if err != nil {
		return
	}

	fallback := []string{"*"}
	got := EnvSplit("TEST_ENV_SPLIT_MISSING", ",", fallback)
	if len(got) != 1 || got[0] != "*" {
		t.Errorf("EnvSplit() = %v, want [*]", got)
	}
}
