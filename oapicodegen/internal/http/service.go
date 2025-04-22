package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/tuanvumaihuynh/go-playground/oapi-codegen/internal/http/gen"
	"github.com/tuanvumaihuynh/go-playground/oapi-codegen/internal/http/swagger"
)

type Config struct {
	Port          int
	EnableSwagger bool
}

type Service struct {
	cfg Config
	log *slog.Logger
}

type CleanupFunc func(ctx context.Context) error

func New(cfg Config, log *slog.Logger) *Service {
	return &Service{
		cfg: cfg,
		log: log,
	}
}

func (s *Service) Run() (CleanupFunc, error) {
	mux := http.NewServeMux()

	if s.cfg.EnableSwagger {
		swagger.Register(mux)
	}
	s.RegisterHandlers(mux)

	return s.RunWithServer(mux)
}

func (s *Service) RunWithServer(handler http.Handler) (CleanupFunc, error) {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.cfg.Port),
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Error("failed to start HTTP server", slog.Any("error", err))
		}
	}()

	return func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	}, nil
}

func (s *Service) RegisterHandlers(mux *http.ServeMux) {
	handler := s.newHandler()
	strictHandler := gen.NewStrictHandlerWithOptions(
		handler,
		[]gen.StrictMiddlewareFunc{},
		gen.StrictHTTPServerOptions{},
	)
	gen.HandlerWithOptions(strictHandler, gen.StdHTTPServerOptions{
		BaseURL:     "/api/v1",
		BaseRouter:  mux,
		Middlewares: []gen.MiddlewareFunc{},
	})
}

var _ gen.StrictServerInterface = (*handler)(nil)

type handler struct {
	*userHandler
}

func (s *Service) newHandler() *handler {
	return &handler{
		userHandler: newUserHandler(),
	}
}
