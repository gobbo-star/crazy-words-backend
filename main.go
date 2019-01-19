package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}
var rooms map[string]Room

func main() {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	fmt.Println("server is starting")
	rooms = make(map[string]Room, 0)
	rand.Seed(time.Now().UnixNano())

	defer fmt.Println("server is stopped")
	http.HandleFunc("/", serve)
	_ = http.ListenAndServe(":8080", nil)
	_, _ = fmt.Scanln()
}

func serve(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		var resp []byte
		switch string(message) {
		case "CREATE_ROOM":
			roomName := fmt.Sprintf("r-%v", rand.Intn(5))
			room := Room{}
			rooms[roomName] = room
			resp = []byte(roomName)
		case "GET_ROOMS":
			resp, _ = json.Marshal(rooms)
		default:
			resp = []byte("UNKNOWN SIG")
		}
		err = ws.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
