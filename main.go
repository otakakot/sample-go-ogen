package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/otakakot/sample-go-ogen/pkg/api"
)

func main() {
	hdl, err := api.NewServer(&Handler{})
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           hdl,
		ReadHeaderTimeout: 30 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		slog.Info("start server listen")

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ctx.Done()

	slog.Info("start server shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	slog.Info("done server shutdown")
}

var _ api.Handler = (*Handler)(nil)

type Handler struct{}

// PostHealth implements api.Handler.
func (*Handler) PostHealth(ctx context.Context, req *api.HealthRequestSchema) (api.PostHealthRes, error) {
	return &api.HealthResponseSchema{
		Message: req.Message,
	}, nil
}

// GetHealth implements api.Handler.
func (*Handler) GetHealth(ctx context.Context, params api.GetHealthParams) (api.GetHealthRes, error) {
	return &api.HealthResponseSchema{
		Message: params.Message,
	}, nil
}
