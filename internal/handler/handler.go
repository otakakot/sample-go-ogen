package handler

import (
	"context"

	"github.com/otakakot/sample-go-ogen/pkg/api"
)

var _ api.Handler = (*Handler)(nil)

type Handler struct{}

// PostHealth implements api.Handler.
func (*Handler) PostHealth(ctx context.Context, req *api.HealthRequestSchema) (api.PostHealthRes, error) {
	return &api.HealthResponseSchema{
		Message: req.Message,
	}, nil
}

// GetHealth implements api.Handler.
func (*Handler) GetHealth(ctx context.Context, params api.GetHealthParams) (api.GetHealthRes, error) {
	return &api.HealthResponseSchema{
		Message: params.Message,
	}, nil
}

// Test implements api.Handler.
func (h *Handler) Test(ctx context.Context, req *api.TestReq) (api.TestRes, error) {
	switch req.GetStatus() {
	case 200:
		return &api.TestOK{}, nil
	case 201:
		return &api.TestCreated{}, nil
	case 400:
		return &api.TestBadRequest{}, nil
	case 401:
		return &api.TestUnauthorized{}, nil
	case 403:
		return &api.TestForbidden{}, nil
	case 404:
		return &api.TestNotFound{}, nil
	default:
		return &api.TestInternalServerError{}, nil
	}
}
