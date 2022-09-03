package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/driver"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/handlers"
)

const portNumber = ":8001"

var app config.AppConfig

func main() {
	app.Production = false
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=pubquiz user=postgres password=postgres")

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting app on port %v\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
