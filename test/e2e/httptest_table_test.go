package e2e_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/otakakot/sample-go-ogen/internal/handler"
	"github.com/otakakot/sample-go-ogen/internal/security"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func TestHttpTableTest(t *testing.T) {
	t.Parallel()

	hdl, err := api.NewServer(
		&handler.Handler{},
		&security.Security{},
	)
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	type args struct {
		ctx   context.Context
		req   *api.TestReq
		token string
	}

	type want struct {
		status int
		body   any
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "200",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 200,
				},
				token: "token",
			},
			want: want{
				status: http.StatusOK,
				body: api.OKResponseSchema{
					Message: "OK",
				},
			},
		},
		{
			name: "400",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 400,
				},
				token: "token",
			},
			want: want{
				status: http.StatusBadRequest,
				body: api.Error{
					Message: "bad request",
				},
			},
		},
		{
			name: "401",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 200,
				},
				token: "invalid",
			},
			want: want{
				status: http.StatusForbidden,
				body: api.Error{
					Message: "operation Test: security \"BearerAuth\": forbidden",
				},
			},
		},
		{
			name: "500",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 200,
				},
				token: "",
			},
			want: want{
				status: http.StatusInternalServerError,
				body: api.Error{
					Message: "operation Test: security \"\": security requirement is not satisfied",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			body := &bytes.Buffer{}

			if err := json.NewEncoder(body).Encode(tt.args.req); err != nil {
				t.Error(err)

				return
			}

			req, err := http.NewRequest(http.MethodPost, "/test", body)
			if err != nil {
				t.Error(err)

				return
			}

			req.Header.Set("Content-Type", "application/json")

			if tt.args.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.args.token)
			}

			res := httptest.NewRecorder()

			hdl.ServeHTTP(res, req)

			t.Cleanup(func() {
				if err := res.Result().Body.Close(); err != nil {
					t.Error(err)
				}
			})

			if res.Result().StatusCode != tt.want.status {
				t.Errorf("status code = %d, want %d", res.Result().StatusCode, tt.want.status)
			}

			bt := &bytes.Buffer{}

			if err := json.NewEncoder(bt).Encode(tt.want.body); err != nil {
				t.Error(err)

				return
			}

			want := strings.ReplaceAll(bt.String(), "\n", "")

			got := res.Body.String()

			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("response body mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
