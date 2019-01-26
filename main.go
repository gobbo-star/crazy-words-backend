package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}
var runeSet []rune
var room Room

//var participants
func main() {
	startServer()
}

func startServer() {
	fmt.Println("server is starting")
	defer fmt.Println("server is stopped")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	room = NewRoom(15 * time.Millisecond)
	go room.start()

	http.HandleFunc("/", serve)
	_ = http.ListenAndServe(":8080", nil)
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
		case "HINT":
			//resp = []byte(word)
		case "CONNECT":
			//resp = []byte(fmt.Sprintf("new: %v", len(word)))
		case "EXIT":
			break
		default:
			//if word == string(message) {
			//	word = newWord(randLen())
			//	resp = []byte(fmt.Sprintf("%v is right. new: %v", string(message), len(word)))
			//} else {
			//	resp = []byte("wrong")
			//}
		}
		err = ws.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
