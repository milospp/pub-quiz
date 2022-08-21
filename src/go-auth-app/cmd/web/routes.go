package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/milospp/pub-quiz/pkg/config"
	"github.com/milospp/pub-quiz/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(MiddlewareLogger)

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/login", handlers.Repo.Login)
	mux.Post("/register", handlers.Repo.Register)
	mux.Get("/quizzes", handlers.Repo.GetQuizes)

	return mux
}
