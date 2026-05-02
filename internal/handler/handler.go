package handler

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/otakakot/sample-go-ogen/internal/model"
	"github.com/otakakot/sample-go-ogen/internal/usecase"
	"github.com/otakakot/sample-go-ogen/pkg/api"
)

var _ api.Handler = (*Handler)(nil)

type Handler struct {
	uc usecase.Usecase
}

func New(
	uc usecase.Usecase,
) *Handler {
	return &Handler{
		uc: uc,
	}
}

// PostHealth implements api.Handler.
func (hdl *Handler) PostHealth(ctx context.Context, req *api.HealthRequestSchema) (*api.HealthResponseSchema, error) {
	return &api.HealthResponseSchema{
		Message: req.Message,
	}, nil
}

// GetHealth implements api.Handler.
func (hdl *Handler) GetHealth(ctx context.Context, params api.GetHealthParams) (*api.HealthResponseSchema, error) {
	return &api.HealthResponseSchema{
		Message: params.Message,
	}, nil
}

// Test implements api.Handler.
func (hdl *Handler) Test(ctx context.Context, req *api.TestReq) (api.TestRes, error) {
	if req.GetStatus() == http.StatusOK {
		return &api.OKResponseSchema{
			Message: "ok",
		}, nil
	}

	if _, err := hdl.uc.Switch(ctx, usecase.SwitchInput{
		Status: req.Status,
	}); err != nil {
		return nil, err
	}

	return nil, &api.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.Error{
			Message: "internal server error",
		},
	}
}

// NewError implements api.Handler.
func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	slog.ErrorContext(ctx, err.Error())

	switch {
	case errors.Is(err, model.ErrBadRequest):
		return &api.ErrorStatusCode{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: err.Error(),
			},
		}
	case errors.Is(err, model.ErrUnauthorized):
		return &api.ErrorStatusCode{
			StatusCode: http.StatusUnauthorized,
			Response: api.Error{
				Message: err.Error(),
			},
		}
	case errors.Is(err, model.ErrForbidden):
		return &api.ErrorStatusCode{
			StatusCode: http.StatusForbidden,
			Response: api.Error{
				Message: err.Error(),
			},
		}
	case errors.Is(err, model.ErrNotFound):
		return &api.ErrorStatusCode{
			StatusCode: http.StatusNotFound,
			Response: api.Error{
				Message: err.Error(),
			},
		}
	}

	return &api.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.Error{
			Message: err.Error(),
		},
	}
}
