package e2e_test

import (
	"context"
	"os"
	"testing"

	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func TestE2e(t *testing.T) {
	t.Parallel()

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:8080"
	}

	cli, err := api.NewClient(endpoint)
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
