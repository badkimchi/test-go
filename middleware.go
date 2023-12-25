package main

import (
	"github.com/dillonstreator/opentelemetry-go-contrib/instrumentation/net/http/otelhttp"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
	"time"
)

func setMiddleware(mux *chi.Mux, logger *slog.Logger, cfg *config) {
	mux.Use(middleware.Recoverer)
	mux.Use(trustProxy(logger))
	mux.Use(otelhttp.NewMiddleware("chi"))
	mux.Use(requestLogger(logger, cfg))
	mux.Use(allowCORS())
}

func requestLogger(logger *slog.Logger, cfg *config) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()

				traceID := trace.SpanFromContext(r.Context()).SpanContext().TraceID()
				var reqID string

				if id, err := uuid.Parse(r.Header.Get("x-request-id")); err == nil {
					reqID = id.String()
				} else {
					reqID = uuid.NewString()
				}

				l := logger.With("reqId", reqID, "traceId", traceID)

				ww := middleware.NewWrapResponseWriter(w, 0)
				rc := newByteReadCloser(r.Body)
				r.Body = http.MaxBytesReader(w, rc, cfg.maxAllowedRequestBytes)

				// overwrite `r`'s memory so that recoverer can access the log entry
				*r = *setLogger(r, l)
				*r = *middleware.WithLogEntry(r, newLogEntry(l))

				handler.ServeHTTP(ww, r)

				l.Info(
					"Request handled",
					slog.String("method", r.Method),
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
					slog.String("ua", r.UserAgent()),
					slog.String("ip", r.RemoteAddr),
					slog.Int("bw", ww.BytesWritten()),
					slog.Int64("br", rc.BytesRead()),
					slog.Int("status", ww.Status()),
					slog.Duration("duration", time.Since(start)),
				)
			},
		)
	}
}

func allowCORS() func(http.Handler) http.Handler {
	corsOptions := cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS",
			},
			AllowedHeaders: []string{
				"X-PINGOTHER", "Accept", "Origin", "X-Auth-Token", "Authorization",
				"Content-Type", "X-CSRF-Token", "Cache-Control", "Pragma",
			},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           3600,
		},
	)
	return corsOptions.Handler
}
