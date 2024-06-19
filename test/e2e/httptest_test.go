package e2e_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/otakakot/sample-go-ogen/internal/handler"
	"github.com/otakakot/sample-go-ogen/internal/security"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func TestHttpTest(t *testing.T) {
	t.Parallel()

	hdl, err := api.NewServer(
		&handler.Handler{},
		&security.Security{},
	)
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	t.Run("200", func(t *testing.T) {
		t.Parallel()

		body := &bytes.Buffer{}

		v := &api.TestReq{
			Status: 200,
		}

		if err := json.NewEncoder(body).Encode(v); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/test", body)
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")

		req.Header.Set("Authorization", "Bearer token")

		res := httptest.NewRecorder()

		hdl.ServeHTTP(res, req)

		t.Cleanup(func() {
			if err := res.Result().Body.Close(); err != nil {
				t.Log(err)
			}
		})

		if res.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
		}

		var got api.OKResponseSchema

		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if got.Message != "OK" {
			t.Errorf("expected message %s, got %s", "OK", got.Message)
		}
	})

	t.Run("201", func(t *testing.T) {
		t.Parallel()

		body := &bytes.Buffer{}

		v := &api.TestReq{
			Status: 201,
		}

		if err := json.NewEncoder(body).Encode(v); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/test", body)
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")

		req.Header.Set("Authorization", "Bearer token")

		res := httptest.NewRecorder()

		hdl.ServeHTTP(res, req)

		t.Cleanup(func() {
			if err := res.Result().Body.Close(); err != nil {
				t.Log(err)
			}
		})

		if res.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, res.Code)
		}

		var got api.CreatedResponseSchema

		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if got.Message != "Created" {
			t.Errorf("expected message %s, got %s", "Created", got.Message)
		}
	})

	t.Run("400", func(t *testing.T) {
		t.Parallel()

		body := &bytes.Buffer{}

		v := &api.TestReq{
			Status: 400,
		}

		if err := json.NewEncoder(body).Encode(v); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/test", body)
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")

		req.Header.Set("Authorization", "Bearer token")

		res := httptest.NewRecorder()

		hdl.ServeHTTP(res, req)

		t.Cleanup(func() {
			if err := res.Result().Body.Close(); err != nil {
				t.Log(err)
			}
		})

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.Code)
		}

		var got api.Error

		if err := json.NewDecoder(res.Result().Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if got.GetMessage() != "bad request" {
			t.Errorf("expected message %s, got %s", "bad request", got.GetMessage())
		}

		// MEMO:
		// server では api.ErrorStatusCode で返すが body には api.Error が入る
	})

	t.Run("403", func(t *testing.T) {
		t.Parallel()

		body := &bytes.Buffer{}

		v := &api.TestReq{
			Status: 200,
		}

		if err := json.NewEncoder(body).Encode(v); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/test", body)
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")

		req.Header.Set("Authorization", "Bearer hoge")

		res := httptest.NewRecorder()

		hdl.ServeHTTP(res, req)

		t.Cleanup(func() {
			if err := res.Result().Body.Close(); err != nil {
				t.Log(err)
			}
		})

		if res.Code != http.StatusForbidden {
			t.Errorf("expected status code %d, got %d", http.StatusForbidden, res.Code)
		}
	})

	t.Run("500", func(t *testing.T) {
		t.Parallel()

		body := &bytes.Buffer{}

		v := &api.TestReq{
			Status: 200,
		}

		if err := json.NewEncoder(body).Encode(v); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/test", body)
		if err != nil {
			t.Error(err)
		}

		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		hdl.ServeHTTP(res, req)

		t.Cleanup(func() {
			if err := res.Result().Body.Close(); err != nil {
				t.Log(err)
			}
		})

		if res.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, res.Code)
		}
	})

	t.Run("400", func(t *testing.T) {
		t.Parallel()

		body := &bytes.Buffer{}

		v := &api.TestReq{
			Status: 200,
		}

		if err := json.NewEncoder(body).Encode(v); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/test", body)
		if err != nil {
			t.Error(err)
		}

		// MEMO: Content-Type が設定されていない

		req.Header.Set("Authorization", "Bearer token")

		res := httptest.NewRecorder()

		hdl.ServeHTTP(res, req)

		t.Cleanup(func() {
			if err := res.Result().Body.Close(); err != nil {
				t.Log(err)
			}
		})

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.Code)
		}
	})
}
