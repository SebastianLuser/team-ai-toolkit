package boot

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/educabot/team-ai-toolkit/errors"
	"github.com/gin-gonic/gin"
)

// Server wraps an http.Server with configured timeouts and graceful shutdown.
type Server struct {
	http *http.Server
}

// NewServer creates an HTTP server wrapping a Gin engine.
func NewServer(port string, engine *gin.Engine) *Server {
	return &Server{
		http: &http.Server{
			Addr:         ":" + port,
			Handler:      engine,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

// Run starts the HTTP server. Blocks until the server stops.
func (s *Server) Run() {
	slog.Info("server listening", "addr", s.http.Addr)
	if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server failed", "err", err)
	}
}

// Shutdown gracefully stops the server with a 10-second timeout.
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.http.Shutdown(ctx); err != nil {
		slog.Error("shutdown failed", "err", err)
	}
	slog.Info("server stopped")
}
