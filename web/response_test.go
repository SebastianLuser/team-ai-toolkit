package web

import (
	"net/http"
	"testing"
)

func TestJSON_SetsStatusAndBody(t *testing.T) {
	resp := JSON(http.StatusOK, map[string]string{"key": "val"})
	if resp.Status != http.StatusOK {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusOK)
	}
	if resp.Body == nil {
		t.Error("Body should not be nil")
	}
}

func TestOK_WrapsInData(t *testing.T) {
	resp := OK("hello")
	if resp.Status != http.StatusOK {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusOK)
	}
	body := resp.Body.(map[string]any)
	if body["data"] != "hello" {
		t.Errorf("data = %v, want %q", body["data"], "hello")
	}
}

func TestCreated_Returns201(t *testing.T) {
	resp := Created(map[string]int{"id": 1})
	if resp.Status != http.StatusCreated {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusCreated)
	}
}

func TestNoContent_Returns204(t *testing.T) {
	resp := NoContent()
	if resp.Status != http.StatusNoContent {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusNoContent)
	}
}

func TestErr_SetsErrorStructure(t *testing.T) {
	resp := Err(http.StatusBadRequest, "invalid", "bad input")
	if resp.Status != http.StatusBadRequest {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusBadRequest)
	}
	body := resp.Body.(map[string]any)
	errMap := body["error"].(map[string]string)
	if errMap["code"] != "invalid" {
		t.Errorf("code = %q, want %q", errMap["code"], "invalid")
	}
	if errMap["message"] != "bad input" {
		t.Errorf("message = %q, want %q", errMap["message"], "bad input")
	}
}

func TestPaginated_IncludesMetadata(t *testing.T) {
	resp := Paginated([]string{"a", "b"}, 50, 2, 20)
	if resp.Status != http.StatusOK {
		t.Errorf("Status = %d, want %d", resp.Status, http.StatusOK)
	}
	body := resp.Body.(map[string]any)
	if body["total"] != 50 {
		t.Errorf("total = %v, want 50", body["total"])
	}
	if body["page"] != 2 {
		t.Errorf("page = %v, want 2", body["page"])
	}
	if body["per_page"] != 20 {
		t.Errorf("per_page = %v, want 20", body["per_page"])
	}
}
