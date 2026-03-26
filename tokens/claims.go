package tokens

import "github.com/golang-jwt/jwt/v5"

// Claims represents the JWT payload emitted by the Auth Service.
// All Educabot backends expect this structure in every authenticated request.
type Claims struct {
	UserID int64    `json:"sub"`
	OrgID  int64    `json:"org_id"`
	Roles  []string `json:"roles"`
	Email  string   `json:"email"`
	Name   string   `json:"name"`
	jwt.RegisteredClaims
}

// HasRole checks if the claims contain a specific role.
func (c Claims) HasRole(role string) bool {
	for _, r := range c.Roles {
		if r == role {
			return true
		}
	}
	return false
}

// HasAnyRole checks if the claims contain at least one of the given roles.
func (c Claims) HasAnyRole(roles ...string) bool {
	for _, role := range roles {
		if c.HasRole(role) {
			return true
		}
	}
	return false
}
