package e2e_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/otakakot/sample-go-ogen/internal/handler"
	"github.com/otakakot/sample-go-ogen/internal/model"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	hdl := &handler.Handler{}

	type args struct {
		ctx context.Context
		req *api.TestReq
	}

	tests := []struct {
		name     string
		args     args
		want     api.TestRes
		wantErr  bool
		checkErr func(*testing.T, error)
	}{
		{
			name: "200",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 200,
				},
			},
			want: &api.OKResponseSchema{
				Message: "ok",
			},
			wantErr: false,
			checkErr: func(t *testing.T, err error) {
				t.Helper()
			},
		},
		{
			name: "400",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 400,
				},
			},
			want:    nil,
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				if !errors.Is(err, model.ErrBadRequest) {
					t.Errorf("Handler.Test() error = %v, wantErr %v", err, model.ErrBadRequest)
				}
			},
		},
		{
			name: "401",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 401,
				},
			},
			want:    nil,
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				if !errors.Is(err, model.ErrUnauthorized) {
					t.Errorf("Handler.Test() error = %v, wantErr %v", err, model.ErrUnauthorized)
				}
			},
		},
		{
			name: "403",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 403,
				},
			},
			want:    nil,
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				if !errors.Is(err, model.ErrForbidden) {
					t.Errorf("Handler.Test() error = %v, wantErr %v", err, model.ErrForbidden)
				}
			},
		},
		{
			name: "404",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 404,
				},
			},
			want:    nil,
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				if !errors.Is(err, model.ErrNotFound) {
					t.Errorf("Handler.Test() error = %v, wantErr %v", err, model.ErrNotFound)
				}
			},
		},
		{
			name: "500",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{
					Status: 500,
				},
			},
			want:    nil,
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				if !errors.Is(err, model.ErrInternalServer) {
					t.Errorf("Handler.Test() error = %v, wantErr %v", err, model.ErrInternalServer)
				}
			},
		},
		{
			name: "500",
			args: args{
				ctx: context.Background(),
				req: &api.TestReq{},
			},
			want:    nil,
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				var target *api.ErrorStatusCode

				if !errors.As(err, &target) {
					t.Errorf("Handler.Test() error = %v, wantErr %v", err, api.ErrorStatusCode{})
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt // ref: https://github.com/golang/go/issues/66876
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := hdl.Test(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.Test() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.Test() = %v, want %v", got, tt.want)
			}
			if !tt.wantErr {
				return
			}
			tt.checkErr(t, err)
		})
	}
}
