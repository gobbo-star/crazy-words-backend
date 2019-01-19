package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{}
var word string
var runeSet []rune

func main() {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	fmt.Println("server is starting")
	defer fmt.Println("server is stopped")
	rand.Seed(time.Now().UnixNano())
	genWordsPool()
	word = newWord(randLen())
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
		case "HINT":
			resp = []byte(word)
		case "CONNECT":
			resp = []byte(fmt.Sprintf("new: %v", len(word)))
		case "EXIT":
			break
		default:
			if word == string(message) {
				word = newWord(randLen())
				resp = []byte(fmt.Sprintf("%v is right. new: %v", string(message), len(word)))
			} else {
				resp = []byte("wrong")
			}
		}
		err = ws.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func randLen() int {
	return rand.Intn(5) + 3
}

func newWord(wLen int) string {
	w := strings.Builder{}
	for ; wLen > 0; wLen-- {
		w.WriteRune(runeSet[rand.Intn(len(runeSet))])
	}
	return w.String()
}

func genWordsPool() {
	for r := 'a'; r <= 'z'; r++ {
		runeSet = append(runeSet, r)
	}
}
