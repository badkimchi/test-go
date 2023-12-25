package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func defineRoutes(mux *chi.Mux, cfg *config) {
	handlers, err := controllers()
	if err != nil {
		panic(err)
	}
	defineStaticRoutes(mux)
	mux.Get(cfg.healthEndpoint, handleHealthCheck)
	mux.Get("/test", handlers.AuthC.TestGet)
}

func defineStaticRoutes(mux *chi.Mux) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	workDir := filepath.Dir(ex)
	feDir := "frontend/dist"
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

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
