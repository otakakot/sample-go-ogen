package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/otakakot/sample-go-ogen/internal/handler"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func TestHttpTest(t *testing.T) {
	t.Parallel()

	hdl, err := api.NewServer(&handler.Handler{})
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	t.Run("GET /health", func(t *testing.T) {
		t.Parallel()

		want := "test"

		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		query := req.URL.Query()
		query.Add("message", want)
		req.URL.RawQuery = query.Encode()

		res := httptest.NewRecorder()

		hdl.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("unexpected status code: got=%d, want=%d", res.Code, http.StatusOK)
		}

		defer func() {
			if err := res.Result().Body.Close(); err != nil {
				t.Error(err)
			}
		}()

		var got *api.HealthResponseSchema

		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if got.Message != want {
			t.Errorf("unexpected message: got=%q, want=%q", got.Message, want)
		}
	})

	t.Run("POST /health", func(t *testing.T) {
		t.Parallel()

		want := "test"

		v := api.HealthRequestSchema{
			Message: want,
		}

		body := &bytes.Buffer{}

		if err := json.NewEncoder(body).Encode(v); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/health", body)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json") // Must set Content-Type

		res := httptest.NewRecorder()

		hdl.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("unexpected status code: got=%d, want=%d", res.Code, http.StatusOK)
		}

		defer func() {
			if err := res.Result().Body.Close(); err != nil {
				t.Error(err)
			}
		}()

		var got *api.HealthResponseSchema

		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if got.Message != want {
			t.Errorf("unexpected message: got=%q, want=%q", got.Message, want)
		}
	})
}
