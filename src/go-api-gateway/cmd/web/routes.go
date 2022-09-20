package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/milospp/pub-quiz/src/go-api-gateway/internal/config"
	"github.com/milospp/pub-quiz/src/go-api-gateway/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(CorsMiddleware)

	r.Get("/", handlers.Repo.RedirectToAuth)
	r.Post("/login", handlers.Repo.RedirectToAuth)
	r.Post("/register", handlers.Repo.RedirectToAuth)
	r.Get("/profile", handlers.Repo.RedirectToAuth)

	r.Get("/quizzes", handlers.Repo.RedirectToQuiz)
	r.Get("/quiz/{code}", handlers.Repo.RedirectToQuiz)
	r.Post("/quiz", handlers.Repo.RedirectToQuiz)

	r.Get("/quiz-stats/{quiz_id}", handlers.Repo.RedirectToStats)

	return r
}
