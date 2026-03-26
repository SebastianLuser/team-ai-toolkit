package pagination

import (
	"strconv"

	"github.com/educabot/team-ai-toolkit/web"
)

const (
	defaultPage    = 1
	defaultPerPage = 20
	maxPerPage     = 100
)

// Pagination holds paging parameters parsed from query string.
type Pagination struct {
	Page    int
	PerPage int
}

// Offset returns the SQL offset for the current page.
func (p Pagination) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// Limit returns the SQL limit (same as PerPage).
func (p Pagination) Limit() int {
	return p.PerPage
}

// ParseFromQuery extracts page and per_page from query parameters.
// Clamps values to sensible defaults and maximums.
func ParseFromQuery(req web.Request) Pagination {
	page := parseIntOr(req.Query("page"), defaultPage)
	perPage := parseIntOr(req.Query("per_page"), defaultPerPage)

	if page < 1 {
		page = defaultPage
	}
	if perPage < 1 {
		perPage = defaultPerPage
	}
	if perPage > maxPerPage {
		perPage = maxPerPage
	}

	return Pagination{Page: page, PerPage: perPage}
}

func parseIntOr(s string, fallback int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return v
}
