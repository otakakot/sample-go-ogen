package security

import (
	"context"
	"errors"

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
	if t.GetToken() == "" {
		return nil, errors.New("missing token")
	}

	return ctx, nil
}
