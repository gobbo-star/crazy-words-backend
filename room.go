package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"strings"
	"time"
)

type Room struct {
	ticker       *time.Ticker
	W            string
	participants []*Participant
}

func (r *Room) start() {
	for t := range r.ticker.C {
		r.W = newWord(randLen())
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

func (r *Room) guess(bytes []byte) {
	if r.W != string(bytes) {
		return
	}
	r.W = newWord(randLen())
}

func NewRoom(refreshRate time.Duration) *Room {
	r := Room{}
	rand.Seed(time.Now().UnixNano())
	genWordsPool()
	r.W = newWord(randLen())
	r.participants = make([]*Participant, 0)
	r.ticker = time.NewTicker(refreshRate)
	return &r
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
