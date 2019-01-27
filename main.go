package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}
var refresher *time.Ticker
var room *Room
var nameGen *NameGen
var wordGen *WordGen
var colorGen *ColorGen
var animalGen *AnimalGen

func main() {
	startServer()
}

func startServer() {
	fmt.Println("server is starting")
	defer fmt.Println("server is stopped")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	initGens()
	refresher = time.NewTicker(1 * time.Second)
	go func() {
		for range refresher.C {
			room.notify()
		}
	}()
	room = NewRoom(15*time.Second, wordGen)
	go room.start()
	http.HandleFunc("/", serve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initGens() {
	wordsFile := flag.String("words", "SET DEFAULT VALUE to words", "foo")
	colorsFile := flag.String("colors", "SET DEFAULT VALUE to colors", "foo")
	animalsFile := flag.String("animals", "SET DEFAULT VALUE to animals", "foo")
	flag.Parse()
	wordGen = NewWordGen(*wordsFile)
	colorGen = NewColorGen(*colorsFile)
	animalGen = NewAnimalGen(*animalsFile)
	nameGen = NewNameGen(colorGen, animalGen)
}

func serve(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	p := NewParticipant(
		ws,
		nameGen.GenName())
	room.join(p)
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
			suc := room.guess(message)
			if suc {
				p.Score++
			}
		}
		err = ws.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
