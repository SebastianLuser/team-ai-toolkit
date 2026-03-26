package tokens

import (
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/educabot/team-ai-toolkit/web"
)

// NewAuthInterceptor creates a middleware that validates JWT tokens.
// Extracts the Bearer token from the Authorization header, validates it
// with the public key, and injects Claims into the request context.
// Returns 401 if the token is missing or invalid.
func NewAuthInterceptor(publicKey *rsa.PublicKey) web.Interceptor {
	return func(req web.Request) web.Response {
		tokenStr := extractBearerToken(req.Header("Authorization"))
		if tokenStr == "" {
			return web.Err(http.StatusUnauthorized, "missing_token", "authorization header required")
		}

		claims, err := ValidateJWT(tokenStr, publicKey)
		if err != nil {
			return web.Err(http.StatusUnauthorized, "invalid_token", "token is invalid or expired")
		}

		SetClaims(req, claims)
		req.Next()
		return web.Response{}
	}
}

// NewOptionalAuthInterceptor creates a middleware that validates JWT tokens
// if present, but allows unauthenticated requests through.
// If a valid token is found, Claims are injected into context.
func NewOptionalAuthInterceptor(publicKey *rsa.PublicKey) web.Interceptor {
	return func(req web.Request) web.Response {
		tokenStr := extractBearerToken(req.Header("Authorization"))
		if tokenStr != "" {
			claims, err := ValidateJWT(tokenStr, publicKey)
			if err == nil {
				SetClaims(req, claims)
			}
		}
		req.Next()
		return web.Response{}
	}
}

func extractBearerToken(header string) string {
	if header == "" {
		return ""
	}
	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		return ""
	}
	return parts[1]
}
