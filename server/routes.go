package main

import (
	"app/middleware"
	"app/sql/db"
	"github.com/dillonstreator/opentelemetry-go-contrib/instrumentation/net/http/otelhttp"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func defineRoutes(mux *chi.Mux, cfg *Config, logger *slog.Logger, token *jwtauth.JWTAuth, queries *db.Queries) {
	setCommonMiddleware(mux, logger)
	cont, err := controllers(token, queries)
	if err != nil {
		panic(err)
	}
	mux.Get(cfg.HealthEndpoint, handleHealthCheck)
	defineStaticRoutes(mux)
	defineAPIRoutes(cont, mux, token)
}

func setCommonMiddleware(mux *chi.Mux, logger *slog.Logger) {
	mux.Use(chimiddleware.Recoverer)
	mux.Use(middleware.TrustProxy(logger))
	mux.Use(otelhttp.NewMiddleware("chi"))
	mux.Use(middleware.RequestLogger(logger))
	mux.Use(middleware.CorsHeaders())
}

func defineStaticRoutes(mux *chi.Mux) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	workDir := filepath.Dir(ex)
	feDir := "web/dist"
	feBasePath := path.Join(workDir, feDir)
	mux.Get(
		"/assets/*", func(w http.ResponseWriter, r *http.Request) {
			fullFilePath := path.Join(feBasePath, r.URL.Path)
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			http.ServeFile(w, r, fullFilePath)
		},
	)

	mux.Get(
		"/", func(w http.ResponseWriter, r *http.Request) {
			fullFilePath := path.Join(feBasePath, "index.html")
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			http.ServeFile(w, r, fullFilePath)
			//w.Write([]byte(fullFilePath))
		},
	)
}

func defineAPIRoutes(cont reqControllers, mux *chi.Mux, token *jwtauth.JWTAuth) {
	mux.Route("/public", func(public chi.Router) {
		public.Use(middleware.Timeout(time.Second * 4))
		public.Get("/test", cont.AuthC.TestGet)
	})
	mux.Route("/api", func(private chi.Router) {
		private.Use(middleware.Timeout(time.Second * 10))
		private.Use(jwtauth.Verifier(token))
		private.Use(middleware.Authenticator(1))
	})
}

func handleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
