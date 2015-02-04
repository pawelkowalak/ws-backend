package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", ws)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	Subprotocols:    []string{"echo-protocol"},
}

func ws(rw http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Printf("error upgrading HTTP to websocket: %v", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("couldn't read message: %v", err)
			return
		}
		log.Printf("echoing message back. Type: %v, payload: %v", messageType, string(p))
		if err = conn.WriteMessage(messageType, p); err != nil {
			log.Printf("error writing message: %v", err)
			return
		}
	}
}
