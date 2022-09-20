package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/milospp/pub-quiz/src/go-api-gateway/internal/config"
	"github.com/milospp/pub-quiz/src/go-api-gateway/internal/handlers"
)

const portNumber = ":8000"

var app config.AppConfig

func main() {
	app.Production = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting app-gateway on port %v\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
