package applog

import (
	"io"
	"log/slog"
)

// SetupTest configures a silent logger for tests.
func SetupTest() {
	handler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelError})
	slog.SetDefault(slog.New(handler))
}
