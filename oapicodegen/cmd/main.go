package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/tuanvumaihuynh/go-playground/oapi-codegen/internal/http"
	"github.com/tuanvumaihuynh/go-playground/oapi-codegen/pkg/cmdutil"
)

func main() {
	app := &Application{
		cfg: http.Config{
			Port:          3000,
			EnableSwagger: true,
		},
		log: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		ctx: context.Background(),
	}

	interruptChan := cmdutil.InterruptChan()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startHTTPService(app, interruptChan); err != nil {
			app.log.Error("Error starting HTTP service", slog.Any("error", err))
		}
	}()

	wg.Wait()
}

type Application struct {
	cfg http.Config
	log *slog.Logger
	ctx context.Context
}

func startHTTPService(app *Application, interruptChan <-chan any) error {
	service := http.New(app.cfg, app.log)

	cleanup, err := service.Run()
	if err != nil {
		return fmt.Errorf("failed to start http service: %w", err)
	}

	app.log.Info("HTTP service started", slog.Any("config", app.cfg))

	<-interruptChan

	app.log.Info("HTTP service shutting down")

	if err := cleanup(app.ctx); err != nil {
		return fmt.Errorf("failed to shutdown http service: %w", err)
	}
	app.log.Info("HTTP service shutdown complete")

	return nil
}
