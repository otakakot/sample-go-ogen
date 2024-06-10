// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeGetHealthResponse(response GetHealthRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *HealthResponseSchema:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ErrorResponseSchema:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostHealthResponse(response PostHealthRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *HealthResponseSchema:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ErrorResponseSchema:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeTestResponse(response TestRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TestOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *TestCreated:
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))

		return nil

	case *TestBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *TestUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *TestForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		return nil

	case *TestNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *TestInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
