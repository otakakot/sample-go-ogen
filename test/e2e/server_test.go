package e2e_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/otakakot/sample-go-ogen/internal/handler"
	"github.com/otakakot/sample-go-ogen/internal/security"
	"github.com/otakakot/sample-go-ogen/pkg/api"
	"github.com/otakakot/sample-go-ogen/pkg/testx"
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

	t.Cleanup(func() {
		srv.Close()
	})

	sec := &testx.Security{
		T: t,
	}

	nosec := &testx.NoSecurity{
		T: t,
	}

	not := &testx.NoToken{
		T: t,
	}

	t.Run("200", func(t *testing.T) {
		t.Parallel()

		cli, err := api.NewClient(srv.URL, sec)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()

		res, err := cli.Test(ctx, &api.TestReq{
			Status: 200,
		})
		if err != nil {
			t.Fatal(err)
		}

		got, ok := res.(*api.OKResponseSchema)
		if !ok {
			t.Fatalf("unexpected response type: got=%T, want=%T", res, &api.OKResponseSchema{})
		}

		if got.Message != "ok" {
			t.Errorf("unexpected message: got=%q, want=%q", got.Message, "ok")
		}
	})

	t.Run("400", func(t *testing.T) {
		t.Parallel()

		cli, err := api.NewClient(srv.URL, sec)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()

		res, err := cli.Test(ctx, &api.TestReq{
			Status: 400,
		})
		if err != nil {
			t.Fatal(err)
		}

		got, ok := res.(*api.TestBadRequest)
		if !ok {
			t.Fatalf("unexpected response type: got=%T, want=%T", res, &api.TestBadRequest{})
		}

		if got.Message != "bad request" {
			t.Errorf("unexpected message: got=%q, want=%q", got.Message, "bad request")
		}
	})

	t.Run("403", func(t *testing.T) {
		t.Parallel()

		cli, err := api.NewClient(srv.URL, nosec)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()

		res, err := cli.Test(ctx, &api.TestReq{
			Status: 200,
		})
		if err != nil {
			t.Fatal(err)
		}

		got, ok := res.(*api.TestForbidden)
		if !ok {
			t.Fatalf("unexpected response type: got=%T, want=%T", res, &api.TestForbidden{})
		}

		if got.Message != "operation Test: security \"BearerAuth\": forbidden" {
			t.Errorf("unexpected message: got=%q, want=%q", got.Message, "forbidden")
		}
	})

	t.Run("500", func(t *testing.T) {
		t.Parallel()

		cli, err := api.NewClient(srv.URL, not)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()

		_, err = cli.Test(ctx, &api.TestReq{
			Status: 200,
		})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
