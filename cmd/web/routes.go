package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"mybookings.com/config"
	"mybookings.com/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
