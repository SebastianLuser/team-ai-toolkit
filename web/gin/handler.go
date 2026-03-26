package gin

import (
	"github.com/educabot/team-ai-toolkit/web"
	"github.com/gin-gonic/gin"
)

// Adapt converts a framework-agnostic web.Handler into a gin.HandlerFunc.
func Adapt(h web.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := NewRequest(c)
		resp := h(req)
		if resp.Body != nil {
			c.JSON(resp.Status, resp.Body)
		} else {
			c.Status(resp.Status)
		}
	}
}

// AdaptMiddleware converts a framework-agnostic web.Interceptor into a gin.HandlerFunc.
// If the interceptor returns a status >= 400, the chain is aborted.
func AdaptMiddleware(m web.Interceptor) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := NewRequest(c)
		resp := m(req)
		if resp.Status >= 400 {
			c.AbortWithStatusJSON(resp.Status, resp.Body)
			return
		}
		c.Next()
	}
}
