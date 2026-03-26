package pagination

import "testing"

func TestPagination_Offset(t *testing.T) {
	tests := []struct {
		page    int
		perPage int
		want    int
	}{
		{1, 20, 0},
		{2, 20, 20},
		{3, 10, 20},
		{5, 25, 100},
	}

	for _, tt := range tests {
		p := Pagination{Page: tt.page, PerPage: tt.perPage}
		if got := p.Offset(); got != tt.want {
			t.Errorf("Pagination{%d,%d}.Offset() = %d, want %d", tt.page, tt.perPage, got, tt.want)
		}
	}
}

func TestPagination_Limit(t *testing.T) {
	p := Pagination{Page: 1, PerPage: 25}
	if got := p.Limit(); got != 25 {
		t.Errorf("Limit() = %d, want 25", got)
	}
}
