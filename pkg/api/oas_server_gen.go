// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetHealth implements getHealth operation.
	//
	// Health.
	//
	// GET /health
	GetHealth(ctx context.Context, params GetHealthParams) (*HealthResponseSchema, error)
	// PostHealth implements postHealth operation.
	//
	// Health.
	//
	// POST /health
	PostHealth(ctx context.Context, req *HealthRequestSchema) (*HealthResponseSchema, error)
	// Test implements test operation.
	//
	// Test.
	//
	// POST /test
	Test(ctx context.Context, req *TestReq) (TestRes, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
