package security

import (
	"context"
	"log/slog"

	"github.com/otakakot/sample-go-ogen/internal/model"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

var _ api.SecurityHandler = (*Security)(nil)

type Security struct{}

// HandleBearerAuth implements api.SecurityHandler.
func (s *Security) HandleBearerAuth(
	ctx context.Context,
	operationName string,
	t api.BearerAuth,
) (context.Context, error) {
	if t.GetToken() != "token" {
		slog.WarnContext(ctx, "invalid token")

		return nil, model.ErrForbidden
	}

	return ctx, nil
}
