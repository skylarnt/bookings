package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/skylarnt/bookings/pkg/config"
	"github.com/skylarnt/bookings/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
