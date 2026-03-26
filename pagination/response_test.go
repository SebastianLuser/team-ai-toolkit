package pagination

import "testing"

func TestNewResponse_NilDataBecomesEmptySlice(t *testing.T) {
	resp := NewResponse[string](nil, 0, Pagination{Page: 1, PerPage: 20})
	if resp.Data == nil {
		t.Error("Data should be empty slice, not nil")
	}
	if len(resp.Data) != 0 {
		t.Errorf("Data length = %d, want 0", len(resp.Data))
	}
}

func TestNewResponse_PreservesMetadata(t *testing.T) {
	p := Pagination{Page: 3, PerPage: 10}
	resp := NewResponse([]string{"a", "b"}, 42, p)

	if resp.Total != 42 {
		t.Errorf("Total = %d, want 42", resp.Total)
	}
	if resp.Page != 3 {
		t.Errorf("Page = %d, want 3", resp.Page)
	}
	if resp.PerPage != 10 {
		t.Errorf("PerPage = %d, want 10", resp.PerPage)
	}
	if len(resp.Data) != 2 {
		t.Errorf("Data length = %d, want 2", len(resp.Data))
	}
}
