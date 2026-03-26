package web

import "context"

// MockRequest is a test double for the Request interface.
// Set the fields you need per test.
type MockRequest struct {
	Params  map[string]string
	Queries map[string]string
	Headers map[string]string
	BindFn  func(dest any) error
	Values  map[string]any
	Ctx     context.Context
	NextFn  func()
}

func NewMockRequest() *MockRequest {
	return &MockRequest{
		Params:  make(map[string]string),
		Queries: make(map[string]string),
		Headers: make(map[string]string),
		Values:  make(map[string]any),
		Ctx:     context.Background(),
	}
}

func (m *MockRequest) Param(key string) string  { return m.Params[key] }
func (m *MockRequest) Query(key string) string  { return m.Queries[key] }
func (m *MockRequest) Header(key string) string { return m.Headers[key] }
func (m *MockRequest) Context() context.Context { return m.Ctx }

func (m *MockRequest) Bind(dest any) error {
	if m.BindFn != nil {
		return m.BindFn(dest)
	}
	return nil
}

func (m *MockRequest) Set(key string, value any) {
	m.Values[key] = value
}

func (m *MockRequest) Get(key string) (any, bool) {
	v, ok := m.Values[key]
	return v, ok
}

func (m *MockRequest) Next() {
	if m.NextFn != nil {
		m.NextFn()
	}
}

var _ Request = (*MockRequest)(nil)
