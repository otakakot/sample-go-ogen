package usecase

import (
	"context"
	"net/http"

	"github.com/otakakot/sample-go-ogen/internal/model"
)

type Usecase struct{}

func NewUsecase() *Usecase {
	return &Usecase{}
}

type SwitchInput struct {
	Status int
}

type SwitchOutput struct{}

func (uc *Usecase) Switch(
	ctx context.Context,
	input SwitchInput,
) (*SwitchOutput, error) {
	if input.Status < 300 {
		return &SwitchOutput{}, nil
	}

	switch input.Status {
	case http.StatusBadRequest:
		return nil, model.ErrBadRequest
	case http.StatusUnauthorized:
		return nil, model.ErrUnauthorized
	case http.StatusForbidden:
		return nil, model.ErrForbidden
	case http.StatusNotFound:
		return nil, model.ErrNotFound
	}

	return nil, model.ErrInternalServer
}
