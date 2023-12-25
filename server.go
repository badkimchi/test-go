package main

import (
	"context"
	"fmt"
	"github.com/NYTimes/gziphandler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func getServer(logger *slog.Logger, cfg *config) *http.Server {
	mux := chi.NewMux()
	setMiddleware(mux, logger, cfg)
	defineRoutes(mux, cfg)
	muxWithGzip := gziphandler.GzipHandler(mux)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: muxWithGzip,
	}
	return srv
}

type byteReadCloser struct {
	rc io.ReadCloser
	n  int64
}

func newByteReadCloser(r io.ReadCloser) *byteReadCloser {
	return &byteReadCloser{r, 0}
}

func (br *byteReadCloser) Read(p []byte) (int, error) {
	n, err := br.rc.Read(p)
	br.n += int64(n)
	return n, err
}

func (br *byteReadCloser) Close() error {
	return br.rc.Close()
}

func (br *byteReadCloser) BytesRead() int64 {
	return br.n
}

type ctxKey string

const (
	ctxKeyLogger ctxKey = "logger"
)

func getLogger(r *http.Request) *slog.Logger {
	return r.Context().Value(ctxKeyLogger).(*slog.Logger)
}

func setLogger(r *http.Request, l *slog.Logger) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), ctxKeyLogger, l))
}

type logEntry struct {
	logger *slog.Logger
}

var _ middleware.LogEntry = (*logEntry)(nil)

func newLogEntry(logger *slog.Logger) *logEntry {
	return &logEntry{logger}
}

func (l *logEntry) Panic(v interface{}, stack []byte) {
	l.logger.Error("panic caught", slog.Any("panic", v), slog.String("stack", string(stack)))
}

func (l *logEntry) Write(status int, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
}
