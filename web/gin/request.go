package gin

import (
	"context"

	"github.com/educabot/team-ai-toolkit/web"
	"github.com/gin-gonic/gin"
)

// Request adapts *gin.Context to the web.Request interface.
type Request struct {
	ctx *gin.Context
}

func NewRequest(c *gin.Context) *Request {
	return &Request{ctx: c}
}

func (r *Request) Param(key string) string   { return r.ctx.Param(key) }
func (r *Request) Query(key string) string   { return r.ctx.Query(key) }
func (r *Request) Header(key string) string  { return r.ctx.GetHeader(key) }
func (r *Request) Bind(dest any) error       { return r.ctx.ShouldBindJSON(dest) }
func (r *Request) Set(key string, value any) { r.ctx.Set(key, value) }
func (r *Request) Context() context.Context  { return r.ctx.Request.Context() }
func (r *Request) Next()                     { r.ctx.Next() }

func (r *Request) Get(key string) (any, bool) {
	return r.ctx.Get(key)
}

// compile-time check
var _ web.Request = (*Request)(nil)
