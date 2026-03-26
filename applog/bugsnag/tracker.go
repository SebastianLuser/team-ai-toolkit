package bugsnag

import (
	"context"
	"net/http"

	bugsnaggin "github.com/bugsnag/bugsnag-go-gin"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/educabot/team-ai-toolkit/applog"
	"github.com/gin-gonic/gin"
)

// Tracker implements applog.ErrorTracker using Bugsnag.
type Tracker struct {
	engine *gin.Engine
}

// NewTracker initializes Bugsnag and returns a Tracker.
// If apiKey is empty, returns a NoopTracker instead.
func NewTracker(apiKey, env string) applog.ErrorTracker {
	if apiKey == "" {
		return applog.NewNoopTracker()
	}

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       apiKey,
		ReleaseStage: env,
	})

	return &Tracker{}
}

func (t *Tracker) Notify(err error, ctx context.Context) {
	err = bugsnag.Notify(err, ctx)
	if err != nil {
		return
	}
}

func (t *Tracker) Handler(next http.Handler) http.Handler {
	return bugsnag.Handler(next)
}

// GinMiddleware returns a gin.HandlerFunc that integrates Bugsnag with Gin.
// Use this in boot/gin.go when setting up the engine.
func GinMiddleware() gin.HandlerFunc {
	return bugsnaggin.AutoNotify()
}
