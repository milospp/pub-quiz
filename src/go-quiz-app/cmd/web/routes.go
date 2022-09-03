package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(MiddlewareLogger)
	r.Use(CorsMiddleware)
	r.Use(AuthMiddleware)

	r.Get("/quizzes", handlers.Repo.GetQuizes)
	r.Get("/quiz/{code}", handlers.Repo.GetQuiz)
	r.Post("/quiz", handlers.Repo.CreateQuiz)

	return r
}
