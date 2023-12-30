package main

import (
	"app/middleware"
	"app/sql/db"
	"app/util"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/NYTimes/gziphandler"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
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

	ctx := context.Background()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	q := db.New(conn)

	fmt.Println("22222")
	user, err := q.GetAuthor(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Bio)

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
