package boot

import (
	"log/slog"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewEngine creates a Gin engine with standard middleware.
// Configures: recovery, CORS, request logging, and a /health endpoint.
func NewEngine(env string, allowedOrigins []string) *gin.Engine {
	if env == "prod" || env == "staging" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	// Recovery from panics
	engine.Use(gin.Recovery())

	// CORS
	engine.Use(corsMiddleware(allowedOrigins))

	// Request logging with slog
	engine.Use(slogMiddleware())

	// Health check (no auth required)
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return engine
}

func corsMiddleware(allowedOrigins []string) gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           10 * time.Minute,
	}

	if len(allowedOrigins) == 1 && allowedOrigins[0] == "*" {
		config.AllowAllOrigins = true
		config.AllowCredentials = false
	} else {
		config.AllowOrigins = allowedOrigins
	}

	return cors.New(config)
}

func slogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		slog.Info("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", time.Since(start).String(),
		)
	}
}
