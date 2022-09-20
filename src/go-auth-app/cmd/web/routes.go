package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(MiddlewareLogger)
	r.Use(CorsMiddleware)

	r.Get("/", handlers.Repo.Home)
	r.Post("/login", handlers.Repo.Login)
	r.Post("/register", handlers.Repo.Register)

	r.Get("/users/{id}", handlers.Repo.GetUser)
	r.Get("/anonymous-users/{id}", handlers.Repo.GetAnonymousUser)
	r.Post("/anonymous-users", handlers.Repo.CreateAnonymousUser)

	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/profile", handlers.Repo.GetLoggedUser)

	})

	return r
}
