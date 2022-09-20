package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/driver"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/dto"
	sockethandlers "github.com/milospp/pub-quiz/src/go-socket-app/internal/socket-handlers"
)

var ch = make(chan ConnectionData)
var app config.AppConfig

type ConnectionData struct {
	Message []byte
	Conn    net.Conn
}

func HandleSocketService() {
	http.ListenAndServe(":8002", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("NEW CONNECTION")

		go func() {
			defer conn.Close()

			roomId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(roomId)

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Println("error loading")

					msgStruct := dto.SocketRequest{
						RoomID: roomId,
						Method: "DISCONNECTED",
						
					}
					msgStructByte, _ := json.Marshal(msgStruct)

					ch <- ConnectionData{
						Conn:    conn,
						Message: msgStructByte,
					}
					return
				}

				if op == ws.OpText {
					ch <- ConnectionData{
						Message: msg,
						Conn:    conn,
					}
				}
			}
		}()
	}))
}

func HandleLogicServices() {
	fmt.Println("Starting GO routine for handling services")
	app.Production = false
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=pubquiz user=postgres password=postgres")

	if err != nil {
		log.Fatal("cannot log to DB")
	}

	repo := sockethandlers.NewRepo(&app, db)
	sockethandlers.NewHandlers(repo)

	for v := range ch {
		var req dto.SocketRequest
		err := json.Unmarshal(v.Message, &req)
		if err != nil {
			continue
		}
		sockethandlers.Repo.ProcessMessage(req, v.Conn)
		// fmt.Printf("MSG: %v", req)
	}

	fmt.Println("-----END SERVICES------")

}
