package main

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		// AllowedOrigins: []string{"http://localhost:3000"},
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})(next)
}
