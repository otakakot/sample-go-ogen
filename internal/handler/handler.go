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
