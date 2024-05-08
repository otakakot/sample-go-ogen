package e2e_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/otakakot/sample-go-ogen/internal/handler"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func TestE2e(t *testing.T) {
	t.Parallel()

	hdl, err := api.NewServer(&handler.Handler{})
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	srv := httptest.NewServer(hdl)

	cli, err := api.NewClient(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	t.Run("GetHealth", func(t *testing.T) {
		if _, err := cli.GetHealth(ctx, api.GetHealthParams{}); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("PostHealth", func(t *testing.T) {
		if _, err := cli.PostHealth(ctx, &api.HealthRequestSchema{}); err != nil {
			t.Fatal(err)
		}
	})
}
