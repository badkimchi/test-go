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

	mux.Get(
		cfg.healthEndpoint, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	workDir := filepath.Dir(ex)
	basePath := path.Join(workDir, "abcd/dist")
	mux.Get(
		"/assets/*", func(w http.ResponseWriter, r *http.Request) {
			fullFilePath := path.Join(basePath, r.URL.Path)
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			http.ServeFile(w, r, fullFilePath)
		},
	)

	mux.Get(
		"/", func(w http.ResponseWriter, r *http.Request) {
			fullFilePath := path.Join(basePath, "index.html")
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			http.ServeFile(w, r, fullFilePath)
			//w.Write([]byte(fullFilePath))
		},
	)

	mux.Get("/test", handlers.AuthC.TestGet)
}
