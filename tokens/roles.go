package tokens

import (
	"net/http"

	"github.com/educabot/team-ai-toolkit/web"
)

// RequireRole creates a middleware that checks if the authenticated user
// has at least one of the specified roles. Returns 403 if not.
// Must be used AFTER NewAuthInterceptor.
func RequireRole(roles ...string) web.Interceptor {
	return func(req web.Request) web.Response {
		claims := GetClaims(req)
		if claims == nil {
			return web.Err(http.StatusUnauthorized, "missing_claims", "auth middleware must run before role check")
		}

		if !claims.HasAnyRole(roles...) {
			return web.Err(http.StatusForbidden, "forbidden", "insufficient permissions")
		}

		req.Next()
		return web.Response{}
	}
}
