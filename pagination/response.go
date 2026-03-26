package pagination

// PaginatedResponse wraps a list of items with pagination metadata.
type PaginatedResponse[T any] struct {
	Data    []T `json:"data"`
	Total   int `json:"total"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

// NewResponse creates a PaginatedResponse from items and pagination info.
func NewResponse[T any](data []T, total int, p Pagination) PaginatedResponse[T] {
	if data == nil {
		data = []T{}
	}
	return PaginatedResponse[T]{
		Data:    data,
		Total:   total,
		Page:    p.Page,
		PerPage: p.PerPage,
	}
}
