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
			room.Notify()
		}
	}()
	room = NewRoom(15*time.Second, wordGen, 3)
	go room.Start()
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
	room.Join(p)
	go readPump(ws, p)
}

func readPump(ws *websocket.Conn, p *Participant) {
READ:
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break READ
		}
		var resp []byte
		switch string(message) {
		case "EXIT":
			break READ
		default:
			suc := room.Guess(message)
			if suc {
				p.Score++
				if room.MaxScoreReached(p.Score) {
					room.NewGame()
					p.Wins++
				}
				room.NewWord()
			}
		}
		if ce, ok := err.(*websocket.CloseError); ok {
			switch ce.Code {
			case websocket.CloseNormalClosure,
				websocket.CloseGoingAway,
				websocket.CloseNoStatusReceived:
				break READ
			}
		}
		err = ws.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			log.Println(err)
			break READ
		}
	}
	room.Quit(p)
	err := ws.Close()
	if err != nil {
		log.Println(err)
	}
}
