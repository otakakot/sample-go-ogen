package e2e_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/otakakot/sample-go-ogen/internal/handler"
	"github.com/otakakot/sample-go-ogen/internal/security"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func TestServer(t *testing.T) {
	t.Parallel()

	hdl, err := api.NewServer(
		&handler.Handler{},
		&security.Security{},
	)
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	srv := httptest.NewServer(hdl)

	cli, err := api.NewClient(srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	t.Run("GetHealth", func(t *testing.T) {
		t.Parallel()

		want := "test"

		params := api.GetHealthParams{
			Message: want,
		}

		res, err := cli.GetHealth(ctx, params)
		if err != nil {
			t.Fatal(err)
		}

		got, ok := res.(*api.HealthResponseSchema)
		if !ok {
			t.Fatalf("unexpected response: %T", res)
		}

		if got.Message != want {
			t.Errorf("unexpected message: got=%q, want=%q", got.Message, want)
		}
	})

	t.Run("PostHealth", func(t *testing.T) {
		t.Parallel()

		want := "test"

		req := &api.HealthRequestSchema{
			Message: want,
		}

		res, err := cli.PostHealth(ctx, req)
		if err != nil {
			t.Fatal(err)
		}

		got, ok := res.(*api.HealthResponseSchema)
		if !ok {
			t.Fatalf("unexpected response: %T", res)
		}

		if got.Message != want {
			t.Errorf("unexpected message: got=%q, want=%q", got.Message, want)
		}
	})
}
