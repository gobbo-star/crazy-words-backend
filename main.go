package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}
var refresher *time.Ticker
var runeSet []rune
var room *Room

func main() {
	startServer()
}

func startServer() {
	fmt.Println("server is starting")
	defer fmt.Println("server is stopped")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	refresher = time.NewTicker(1 * time.Second)
	go func() {
		for range refresher.C {
			room.notify()
		}
	}()
	room = NewRoom(15 * time.Second)
	go room.start()
	http.HandleFunc("/", serve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serve(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	room.join(ws)
	defer room.quit(ws)
	defer ws.Close()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		var resp []byte
		switch string(message) {
		case "HINT":
			resp = []byte(room.W)
		case "EXIT":
			break
		default:
		}
		err = ws.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
