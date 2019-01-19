package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {
	fmt.Println("server is starting")
	defer fmt.Println("server is stopped")
	http.HandleFunc("/", serve)
	_ = http.ListenAndServe(":8080", nil)
	_, _ = fmt.Scanln()
}

func serve(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
	defer ws.Close()
}
