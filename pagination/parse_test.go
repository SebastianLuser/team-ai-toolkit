package pagination

import (
	"testing"

	"github.com/educabot/team-ai-toolkit/web"
)

func TestParseFromQuery_Defaults(t *testing.T) {
	req := web.NewMockRequest()

	p := ParseFromQuery(req)
	if p.Page != 1 {
		t.Errorf("Page = %d, want 1", p.Page)
	}
	if p.PerPage != 20 {
		t.Errorf("PerPage = %d, want 20", p.PerPage)
	}
}

func TestParseFromQuery_CustomValues(t *testing.T) {
	req := web.NewMockRequest()
	req.Queries["page"] = "3"
	req.Queries["per_page"] = "50"

	p := ParseFromQuery(req)
	if p.Page != 3 {
		t.Errorf("Page = %d, want 3", p.Page)
	}
	if p.PerPage != 50 {
		t.Errorf("PerPage = %d, want 50", p.PerPage)
	}
}

func TestParseFromQuery_ClampsMaxPerPage(t *testing.T) {
	req := web.NewMockRequest()
	req.Queries["per_page"] = "999"

	p := ParseFromQuery(req)
	if p.PerPage != 100 {
		t.Errorf("PerPage = %d, want 100 (max)", p.PerPage)
	}
}

func TestParseFromQuery_ClampsNegativeValues(t *testing.T) {
	req := web.NewMockRequest()
	req.Queries["page"] = "-1"
	req.Queries["per_page"] = "0"

	p := ParseFromQuery(req)
	if p.Page != 1 {
		t.Errorf("Page = %d, want 1", p.Page)
	}
	if p.PerPage != 20 {
		t.Errorf("PerPage = %d, want 20 (default)", p.PerPage)
	}
}

func TestParseFromQuery_InvalidStrings(t *testing.T) {
	req := web.NewMockRequest()
	req.Queries["page"] = "abc"
	req.Queries["per_page"] = "xyz"

	p := ParseFromQuery(req)
	if p.Page != 1 {
		t.Errorf("Page = %d, want 1 (default)", p.Page)
	}
	if p.PerPage != 20 {
		t.Errorf("PerPage = %d, want 20 (default)", p.PerPage)
	}
}
