package web

// Handler is a framework-agnostic HTTP handler function.
// All entrypoint handlers implement this signature.
type Handler func(req Request) Response

// Interceptor is a framework-agnostic middleware function.
// Returns a response to short-circuit the chain, or calls req.Next() to continue.
type Interceptor func(req Request) Response
