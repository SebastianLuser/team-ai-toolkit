package tokens

import (
	"net/http"

	"github.com/educabot/team-ai-toolkit/web"
)

const orgIDKey = "org_id"

// NewTenantInterceptor creates a middleware that extracts org_id from
// the already-validated JWT claims and makes it available in context.
// Must be used AFTER NewAuthInterceptor.
func NewTenantInterceptor() web.Interceptor {
	return func(req web.Request) web.Response {
		claims := GetClaims(req)
		if claims == nil {
			return web.Err(http.StatusUnauthorized, "missing_claims", "auth middleware must run before tenant middleware")
		}

		if claims.OrgID == 0 {
			return web.Err(http.StatusForbidden, "missing_org", "user is not assigned to an organization")
		}

		req.Set(orgIDKey, claims.OrgID)
		req.Next()
		return web.Response{}
	}
}
