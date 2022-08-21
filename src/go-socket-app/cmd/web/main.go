package main

import (
	"net"
	"time"

	"github.com/go-stomp/stomp"
)

func doSomethingWith(f ...interface{}) {

}

func main() {
	netConn, err := net.DialTimeout("tcp", "stomp.server.com:61613", 10*time.Second)
	if err != nil {
		return
	}

	stompConn, err := stomp.Connect(netConn)
	if err != nil {
		return
	}

	defer stompConn.Disconnect()

	doSomethingWith(stompConn)
	return
}
