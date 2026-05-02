package testx

import (
	"context"
	"testing"

	"github.com/otakakot/sample-go-ogen/pkg/api"
)

var _ api.SecuritySource = (*Security)(nil)

type Security struct {
	T *testing.T
}

// BearerAuth implements api.SecuritySource.
func (s *Security) BearerAuth(
	ctx context.Context,
	operationName string,
) (api.BearerAuth, error) {
	s.T.Helper()

	return api.BearerAuth{
		Token: "token",
	}, nil
}

var _ api.SecuritySource = (*NoSecurity)(nil)

type NoSecurity struct {
	T *testing.T
}

// BearerAuth implements api.SecuritySource.
func (n *NoSecurity) BearerAuth(ctx context.Context, operationName string) (api.BearerAuth, error) {
	n.T.Helper()

	return api.BearerAuth{
		Token: "no",
	}, nil
}

var _ api.SecuritySource = (*NoToken)(nil)

type NoToken struct {
	T *testing.T
}

// BearerAuth implements api.SecuritySource.
func (n *NoToken) BearerAuth(ctx context.Context, operationName string) (api.BearerAuth, error) {
	n.T.Helper()

	return api.BearerAuth{}, nil
}
