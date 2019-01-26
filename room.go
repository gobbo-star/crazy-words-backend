package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

type Room struct {
	ticker       *time.Ticker
	W            string
	participants []*Participant
	wg           *WordGen
}

func (r *Room) start() {
	for t := range r.ticker.C {
		r.W = r.wg.GenWord()
		fmt.Println(t)
		fmt.Println(r.W)
	}
}

func (r *Room) join(p *Participant) {
	r.participants = append(r.participants, p)
}

func (r *Room) notify() {
	rs, err := json.Marshal(NewRiddle(r))
	for i := 0; i < len(r.participants); i++ {
		p := r.participants[i]
		if err != nil {
			fmt.Println(err)
		}
		p.Notify(rs)
	}
}

func (r *Room) quit(ws *websocket.Conn) {
	// TODO
	//ws.
}

func (r *Room) guess(bytes []byte) bool {
	success := r.W == string(bytes)
	if success {
		r.W = r.wg.GenWord()
	}
	return success
}

func NewRoom(refreshRate time.Duration, gen *WordGen) *Room {
	r := Room{}
	rand.Seed(time.Now().UnixNano())
	genWordsPool()
	r.wg = gen
	r.W = r.wg.GenWord()
	r.participants = make([]*Participant, 0)
	r.ticker = time.NewTicker(refreshRate)
	return &r
}
