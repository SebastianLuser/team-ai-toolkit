package applog

import (
	"log/slog"
	"os"
)

// Setup configures the global slog logger based on the environment.
//   - prod/staging: JSON output to stdout (for Cloud Logging, Datadog, etc.)
//   - local/develop: human-readable text output
func Setup(env string) {
	var handler slog.Handler
	if env == "prod" || env == "staging" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	slog.SetDefault(slog.New(handler))
}
