package main

import (
	"app/middleware"
	"app/util"
	"context"
	"errors"
	"fmt"
	"github.com/NYTimes/gziphandler"
	"github.com/go-chi/chi"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.DbUser)
	fmt.Println(cfg.DbPassword)
	fmt.Println(cfg.DbHost)
	fmt.Println(cfg.DbPort)
	fmt.Println(cfg.DbName)

	logger := util.NewLogger(os.Stdout, cfg.LogLevel)
	otelShutdown, err := SetupOTelSDK(context.Background(), cfg)
	if err != nil {
		logger.Error("Setting up open telemetry", slog.Any("error", err))
		os.Exit(1)
	}

	mux := chi.NewMux()
	middleware.SetMiddleware(mux, logger)
	defineRoutes(mux, cfg)
	muxWithGzip := gziphandler.GzipHandler(mux)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: muxWithGzip,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Server error", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	logger.Info(fmt.Sprintf("Listening for HTTP on Port %d", cfg.Port))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	sig := <-shutdown
	logger.Info("Shutdown signal received", "signal", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		logger.Error("Server shutdown", slog.Any("error", err))
		os.Exit(1)
	}

	err = otelShutdown(ctx)
	if err != nil {
		logger.Error("Open telemetry shutdown", slog.Any("error", err))
		os.Exit(1)
	}
}
