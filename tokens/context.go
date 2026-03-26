package tokens

import "github.com/educabot/team-ai-toolkit/web"

const claimsKey = "auth_claims"

// SetClaims stores the parsed JWT claims in the request context.
func SetClaims(req web.Request, claims *Claims) {
	req.Set(claimsKey, claims)
}

// GetClaims retrieves the JWT claims from the request context.
// Returns nil if no claims are set (unauthenticated request).
func GetClaims(req web.Request) *Claims {
	v, ok := req.Get(claimsKey)
	if !ok {
		return nil
	}
	claims, ok := v.(*Claims)
	if !ok {
		return nil
	}
	return claims
}

// MustClaims retrieves the JWT claims or panics if missing.
// Use only in handlers behind auth middleware (guaranteed to have claims).
func MustClaims(req web.Request) *Claims {
	claims := GetClaims(req)
	if claims == nil {
		panic("claims not in context — auth middleware missing")
	}
	return claims
}

// UserID extracts the user ID from the request claims.
func UserID(req web.Request) int64 {
	return MustClaims(req).UserID
}

// OrgID extracts the organization ID from the request claims.
func OrgID(req web.Request) int64 {
	return MustClaims(req).OrgID
}
