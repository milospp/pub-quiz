package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/milospp/pub-quiz/src/go-auth-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/driver"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/handlers"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/utils"
)

const portNumber = ":8003"

var app config.AppConfig

func main() {
	app.Production = false

	sqlConnStr := utils.GetSqlConnectionString()
	db, err := driver.ConnectSQL(sqlConnStr)

	utils.InitRandom()

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
