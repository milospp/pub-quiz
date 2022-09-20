package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/cors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/config"
)

func MiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tok := strings.Split(r.Header.Get("Authorization"), " ")
		if len(tok) < 2 {
			next.ServeHTTP(w, r)
			return
		}
		a := tok[1]

		token, err := jwt.ParseWithClaims(a, &config.AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("DontHeckMe"), nil
		})
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(403)
			return
		}
		if token.Valid == false {
			fmt.Println("invladi token")
			w.WriteHeader(403)
			return
		}

		// claims := token.Claims.(*config.AuthClaim)
		next.ServeHTTP(w, r)
	})

}
