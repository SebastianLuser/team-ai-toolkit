package config

import (
	"os"
	"strings"
)

// EnvOr returns the value of the environment variable or the fallback if empty.
func EnvOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// MustEnv returns the value of the environment variable or panics if missing.
// Use for required configuration that the app cannot start without.
func MustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("required env var missing: " + key)
	}
	return v
}

// EnvSplit returns the value of the environment variable split by separator.
// Returns fallback if the variable is empty.
func EnvSplit(key, sep string, fallback []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return strings.Split(v, sep)
}
