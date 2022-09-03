package main

import (
	"net"

	"github.com/milospp/pub-quiz/src/go-socket-app/internal/services"
)

var clients = make(map[*net.Conn]bool)

func main() {
	go services.HandleLogicServices()
	services.HandleSocketService()
}
