package e2e_test

import (
	"context"
	"testing"

	"github.com/otakakot/sample-go-ogen/internal/handler"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	hdl := &handler.Handler{}

	ctx := context.Background()

	t.Run("GetHealth", func(t *testing.T) {
		t.Parallel()

		want := "test"

		params := api.GetHealthParams{
			Message: want,
		}

		res, err := hdl.GetHealth(ctx, params)
		if err != nil {
			t.Fatalf("GetHealth failed: %v", err)
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

		res, err := hdl.PostHealth(ctx, req)
		if err != nil {
			t.Fatalf("PostHealth failed: %v", err)
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
